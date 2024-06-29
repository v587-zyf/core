package core

import (
	"core/log/color"
	"fmt"

	"go.uber.org/zap/zapcore"
)

type Level int8

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

func (l Level) Enabled(lvl Level) bool {
	return lvl >= l
}

var DefaultConsoleEncoderConfig = zapcore.EncoderConfig{
	TimeKey:       "ts",
	LevelKey:      "level",
	NameKey:       "logger",
	CallerKey:     "caller",
	FunctionKey:   zapcore.OmitKey,
	MessageKey:    "msg",
	StacktraceKey: "stacktrace",
	LineEnding:    zapcore.DefaultLineEnding,
	EncodeLevel:   defaultEncodeLevel,
	EncodeTime:    zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
	// EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

var DefaultFileEncoderConfig = zapcore.EncoderConfig{
	TimeKey:       "ts",
	LevelKey:      "level",
	NameKey:       "logger",
	CallerKey:     "caller",
	FunctionKey:   zapcore.OmitKey,
	MessageKey:    "msg",
	StacktraceKey: "stacktrace",
	LineEnding:    zapcore.DefaultLineEnding,
	EncodeLevel:   fileEncodeLevel,
	EncodeTime:    zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
	// EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func defaultEncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	str, ok := levelToCapitalColorStrings[l]
	if !ok {
		str = color.Red.Add(fmt.Sprintf("[LEVEL(%d)]", l))
		levelToCapitalColorStrings[l] = str
	}

	enc.AppendString(str)
}

func fileEncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {

	str, ok := levelToCapitalStrings[l]
	if !ok {
		str = fmt.Sprintf("[LEVEL(%d)]", l)
		levelToCapitalColorStrings[l] = str
	}

	enc.AppendString(str)
}

var levelToCapitalColorStrings = map[zapcore.Level]string{
	DebugLevel: color.Magenta.Add("[DEBUG]"),
	InfoLevel:  color.Blue.Add("[INFO]"),
	WarnLevel:  color.Yellow.Add("[WARN]"),
	ErrorLevel: color.Red.Add("[ERROR]"),
	PanicLevel: color.Red.Add("[PANIC]"),
	FatalLevel: color.Red.Add("[FATAL]"),
}

var levelToCapitalStrings = map[zapcore.Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}
