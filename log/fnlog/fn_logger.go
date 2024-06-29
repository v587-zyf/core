package fnlog

import (
	"os"
	"path"
	"strings"
	"thunder/internal/core/log/core"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FnLogger struct {
	config *Config

	fields []zap.Field

	zapLogger *zap.Logger
}

func New(cfg *Config) (*FnLogger, error) {
	logger := &FnLogger{
		config: cfg.defalut(),

		fields: make([]zap.Field, 0),
	}

	logger.fields = []zap.Field{}

	if cfg.GID != "" {
		logger.fields = append(logger.fields, zap.String("gid", cfg.GID))
	}

	if cfg.SID > 0 {
		logger.fields = append(logger.fields, zap.Int32("sid", cfg.SID))
	}

	level := strings.ToUpper(cfg.Type)
	name := strings.ToLower(cfg.Type)

	rotateOpts := []rotatelogs.Option{
		rotatelogs.WithMaxAge(cfg.MaxAge),
		rotatelogs.WithRotationTime(24 * time.Hour),
		rotatelogs.WithRotationSize(cfg.MaxSize),
	}

	if cfg.LinkEnabled {
		rotateOpts = append(rotateOpts, rotatelogs.WithLinkName("./"+name+".log"))
	}

	out, err := rotatelogs.New(path.Join(cfg.Path, name, name+"-%Y-%m-%d.log"), rotateOpts...)
	if err != nil {
		return nil, err
	}
	fileWriteSyncer := zapcore.AddSync(out)

	fileEncoderConfig := core.DefaultFileEncoderConfig
	fileEncoderConfig.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(level)
	}
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)
	fileCore := zapcore.NewCore(fileEncoder, fileWriteSyncer, zapcore.DebugLevel)

	consoleEncoderConfig := core.DefaultConsoleEncoderConfig
	consoleLevel := "\x1b[34m[" + level + "]\x1b[0m"
	consoleEncoderConfig.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(consoleLevel)
	}
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	var logCore zapcore.Core
	if !cfg.ConsoleDisabled && !cfg.FileDisabled {
		logCore = zapcore.NewTee(fileCore, consoleCore)
	} else if !cfg.FileDisabled {
		logCore = zapcore.NewTee(fileCore)
	} else {
		logCore = zapcore.NewTee(consoleCore)
	}
	logger.zapLogger = zap.New(logCore, zap.Fields(logger.fields...))

	return logger, nil
}

func (logger *FnLogger) Info(msg string, fields ...zapcore.Field) {
	logger.zapLogger.Info(msg, fields...)
}

func (logger *FnLogger) Clone() *FnLogger {
	clone := *logger
	zapClone := *logger.zapLogger

	clone.zapLogger = &zapClone
	return &clone
}

func (logger *FnLogger) With(fields ...zap.Field) {
	logger.zapLogger = logger.zapLogger.With(fields...)
}

func (logger *FnLogger) WithOptions(opts ...zap.Option) {
	logger.zapLogger = logger.zapLogger.WithOptions(opts...)
}

func (logger *FnLogger) WithGID(gid string) {
	if logger.config.GID != "" {
		return
	}
	logger.config.GID = gid
	logger.With(zap.String("gid", gid))
}

func (logger *FnLogger) WithSID(sid int32) {
	if logger.config.SID > 0 {
		return
	}
	logger.config.SID = sid
	logger.With(zap.Int32("sid", sid))
}
