package bro_logger

import (
	"context"
	"errors"
	"net"
	"os"
	"sync"
	"time"

	"go.elastic.co/apm/module/apmzap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var logstashWriter *reconnectingWriter

// reconnectingWriter wraps a net.Conn and handles reconnection logic
type reconnectingWriter struct {
	mu        sync.Mutex
	address   string
	conn      net.Conn
	reconnect chan struct{}
}

func newReconnectingWriter(address string) *reconnectingWriter {
	rw := &reconnectingWriter{
		address:   address,
		reconnect: make(chan struct{}, 1),
	}
	go rw.manageConnection()
	return rw
}

func (rw *reconnectingWriter) manageConnection() {
	for {
		if rw.conn == nil {
			conn, err := net.Dial("tcp", rw.address)
			if err == nil {
				rw.mu.Lock()
				rw.conn = conn
				rw.mu.Unlock()
			} else {
				time.Sleep(5 * time.Second) // Retry after 5 seconds
			}
		}
		<-rw.reconnect
		rw.closeConnection()
	}
}

func (rw *reconnectingWriter) closeConnection() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	if rw.conn != nil {
		rw.conn.Close()
		rw.conn = nil
	}
}

func (rw *reconnectingWriter) Write(p []byte) (n int, err error) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	if rw.conn == nil {
		return 0, errors.New("connection to Logstash is not available")
	}
	n, err = rw.conn.Write(p)
	if err != nil {
		rw.triggerReconnect()
	}
	return n, err
}

func (rw *reconnectingWriter) Sync() error {
	return nil // No-op for reconnectingWriter
}

func (rw *reconnectingWriter) triggerReconnect() {
	select {
	case rw.reconnect <- struct{}{}:
	default: // Avoid blocking if a reconnect signal is already pending
	}
}

func init() {
	initLogger("", zapcore.InfoLevel) // Default initialization
}

// Initialize the logger with a Logstash endpoint
func initLogger(logstashAddr string, logLevel zapcore.Level) {
	var cores []zapcore.Core

	// Create JSON encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Add console core for fallback logging
	consoleCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stdout)), logLevel)
	cores = append(cores, consoleCore)

	// Add Logstash core if address is provided
	if logstashAddr != "" {
		logstashWriter = newReconnectingWriter(logstashAddr)
		logstashCore := zapcore.NewCore(encoder, zapcore.AddSync(logstashWriter), logLevel)
		cores = append(cores, logstashCore)
	}

	// Combine cores
	core := zapcore.NewTee(cores...)

	// Initialize logger
	logger = zap.New(core, zap.WrapCore((&apmzap.Core{}).WrapCore), zap.AddCaller(), zap.AddCallerSkip(1))
}

// Expose logger instance for custom usage
func Instance() *zap.Logger {
	return logger
}

// Log levels with context and APM fields
func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Debug(msg, fields...)
}
func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Info(msg, fields...)
}
func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Warn(msg, fields...)
}
func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Error(msg, fields...)
}
func DPanic(ctx context.Context, msg string, fields ...zapcore.Field) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).DPanic(msg, fields...)
}
func Panic(ctx context.Context, msg string, fields ...zapcore.Field) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Panic(msg, fields...)
}
func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Fatal(msg, fields...)
}

// Configure the logger with custom configuration
func Configure(logstashAddr string, logLevel zapcore.Level) {
	initLogger(logstashAddr, logLevel)
}
