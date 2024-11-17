package bro_logger

import (
	"context"
	"net"
	"os"

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
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// Initialize Logstash writer
func initLogstashWriter(logstashAddr string) zapcore.WriteSyncer {
	conn, err := net.Dial("tcp", logstashAddr) // Use "udp" for UDP transport
	if err != nil {
		panic("Failed to connect to Logstash: " + err.Error())
	}
	return zapcore.AddSync(conn)
}

// Extract APM fields from context
func extractApmToLog(ctx context.Context) []zapcore.Field {
	fields := []zapcore.Field{
		zap.Any("env", ctx.Value("env")),
		zap.Any("service", ctx.Value("service")),
		zap.Any("trace_id", ctx.Value("trace_id")),
		zap.Any("span_id", ctx.Value("span_id")),
	}

	return fields
}

// Expose logger instance for custom usage
func Instance() *zap.Logger {
	return logger
}

// Log levels with context and APM fields
func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, extractApmToLog(ctx)...)
	logger.Debug(msg, fields...)
}
func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, extractApmToLog(ctx)...)
	logger.Info(msg, fields...)
}
func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, extractApmToLog(ctx)...)
	logger.Warn(msg, fields...)
}
func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, extractApmToLog(ctx)...)
	logger.Error(msg, fields...)
}
func DPanic(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, extractApmToLog(ctx)...)
	logger.DPanic(msg, fields...)
}
func Panic(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, extractApmToLog(ctx)...)
	logger.Panic(msg, fields...)
}
func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, extractApmToLog(ctx)...)
	logger.Fatal(msg, fields...)
}

// Initialize the logger with custom configuration
func Configure(logstashAddr string, logLevel zapcore.Level) {
	initLogger(logstashAddr, logLevel)
}

// Helper: Fallback writer to log to the console if Logstash is unavailable
func newFallbackWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(zapcore.Lock(zapcore.AddSync(os.Stdout)))
}
