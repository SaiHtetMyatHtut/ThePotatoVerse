package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the interface for logging
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

// zapLogger implements Logger interface using zap
type zapLogger struct {
	logger *zap.SugaredLogger
}

// New creates a new logger
func New() Logger {
	// Create encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugar := logger.Sugar()

	return &zapLogger{
		logger: sugar,
	}
}

// Debug logs a debug message
func (l *zapLogger) Debug(msg string, args ...interface{}) {
	l.logger.Debugw(msg, args...)
}

// Info logs an info message
func (l *zapLogger) Info(msg string, args ...interface{}) {
	l.logger.Infow(msg, args...)
}

// Warn logs a warning message
func (l *zapLogger) Warn(msg string, args ...interface{}) {
	l.logger.Warnw(msg, args...)
}

// Error logs an error message
func (l *zapLogger) Error(msg string, args ...interface{}) {
	l.logger.Errorw(msg, args...)
}

// Fatal logs a fatal message and exits
func (l *zapLogger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatalw(msg, args...)
}
