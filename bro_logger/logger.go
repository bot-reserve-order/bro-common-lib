package bro_logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	if logger == nil {
		logger, _ = zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
	}
}

func extractApmToLog(ctx context.Context) []zapcore.Field {
	fields := []zapcore.Field{
		zap.Any("env", ctx.Value("env")),
		zap.Any("service", ctx.Value("service")),
		zap.Any("trace_id", ctx.Value("trace_id")),
		zap.Any("span_id", ctx.Value("span_id")),
	}

	return fields
}

func Instance() *zap.Logger {
	return logger
}

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
