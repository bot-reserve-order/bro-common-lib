package bro_logger

import (
	"context"
	"net"
	"os"

	"go.elastic.co/apm/module/apmzap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var logstashWriter zapcore.WriteSyncer

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
	consoleCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(zapcore.AddSync(newFallbackWriter()))), logLevel)
	cores = append(cores, consoleCore)

	// Add Logstash core if address is provided
	if logstashAddr != "" {
		logstashWriter = initLogstashWriter(logstashAddr)
		logstashCore := zapcore.NewCore(encoder, logstashWriter, logLevel)
		cores = append(cores, logstashCore)
	}

	// Combine cores
	core := zapcore.NewTee(cores...)

	// Initialize logger
	logger = zap.New(core, zap.WrapCore((&apmzap.Core{}).WrapCore), zap.AddCaller(), zap.AddCallerSkip(1))
}

// Initialize Logstash writer
func initLogstashWriter(logstashAddr string) zapcore.WriteSyncer {
	conn, err := net.Dial("tcp", logstashAddr) // Use "udp" for UDP transport
	if err != nil {
		panic("Failed to connect to Logstash: " + err.Error())
	}
	return zapcore.AddSync(conn)
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

// Initialize the logger with custom configuration
func Configure(logstashAddr string, logLevel zapcore.Level) {
	initLogger(logstashAddr, logLevel)
}

// Helper: Fallback writer to log to the console if Logstash is unavailable
func newFallbackWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(zapcore.Lock(zapcore.AddSync(os.Stdout)))
}
