package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.StacktraceKey = "tracerKey"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Info wrap for log.info
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug wrap for log.Debug
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error wrap for log.Error
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// Fatal wrap for log.Error
func Fatal(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// Warn wrap for log.Warn
func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

// Sync wrap for log.Sync
func Sync() error {
	return log.Sync()
}
