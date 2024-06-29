package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defLogger *Logger

var defConfig = Config{
	GID: "bar",

	Fields: make([]zap.Field, 0),

	LinkEnabled: true,

	ConsoleDisabled: false,
	FileDisabled:    false,

	CallerSkip: 3,

	SysLevel: zapcore.DebugLevel,

	Path: "./log",

	SysMaxAge: 30 * 24 * time.Hour,

	MaxSize: 300 * 1024 * 1024,
}

func GetDefaultLogger() *Logger {
	return defLogger
}

func Init(config ...*Config) error {
	var cfg *Config
	if len(config) <= 0 {
		cfg = &defConfig
	} else {
		cfg = config[0]
	}
	var err error
	defLogger, err = New(cfg)
	return err
}

func Info(msg string, fields ...zapcore.Field) {
	defLogger.Info(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	defLogger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	defLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	defLogger.Error(msg, fields...)
}

func Crash(msg string, fields ...zapcore.Field) {
	defLogger.Crash(msg, fields...)
}

func Clone() *Logger {
	return defLogger.Clone()
}

func With(fields ...zap.Field) {
	defLogger.With(fields...)
}
