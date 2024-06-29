package syslog

import (
	"core/log/core"
	"os"
	"path"
	"time"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SysLogger struct {
	config *Config

	fields []zap.Field

	infoLogger  *zap.Logger
	errLogger   *zap.Logger
	crashLogger *zap.Logger
}

func New(cfg *Config) (*SysLogger, error) {
	logger := &SysLogger{
		config: cfg.get(),

		fields: make([]zap.Field, 0),
	}

	logger.fields = []zap.Field{}

	if cfg.GID != "" {
		logger.fields = append(logger.fields, zap.String("gid", cfg.GID))
	}

	if cfg.SID > 0 {
		logger.fields = append(logger.fields, zap.Int32("sid", cfg.SID))
	}

	infoOpts := []rotateLogs.Option{
		rotateLogs.WithMaxAge(cfg.MaxAge),
		rotateLogs.WithRotationTime(24 * time.Hour),
		rotateLogs.WithRotationSize(cfg.MaxSize),
	}

	if cfg.LinkEnabled {
		infoOpts = append(infoOpts, rotateLogs.WithLinkName("./info.log"))
	}

	out, err := rotateLogs.New(path.Join(cfg.Path, "info-%Y-%m-%d.log"), infoOpts...)
	if err != nil {
		return nil, err
	}
	infoFileWriteSyncer := zapcore.AddSync(out)

	errOpts := []rotateLogs.Option{
		rotateLogs.WithMaxAge(cfg.MaxAge),
		rotateLogs.WithRotationTime(24 * time.Hour),
		rotateLogs.WithRotationSize(cfg.MaxSize),
	}

	if cfg.LinkEnabled {
		errOpts = append(errOpts, rotateLogs.WithLinkName("./error.log"))
	}

	out, err = rotateLogs.New(path.Join(cfg.Path, "error-%Y-%m-%d.log"), errOpts...)
	if err != nil {
		return nil, err
	}
	errFileWriteSyncer := zapcore.AddSync(out)

	// info log
	fileEncoder := zapcore.NewJSONEncoder(core.DefaultFileEncoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(core.DefaultConsoleEncoderConfig)
	infoFileCore := zapcore.NewCore(fileEncoder, infoFileWriteSyncer, zap.DebugLevel)
	infoConsoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)

	var logCore zapcore.Core
	if !cfg.ConsoleDisabled && !cfg.FileDisabled {
		logCore = zapcore.NewTee(infoFileCore, infoConsoleCore)
	} else if !cfg.FileDisabled {
		logCore = zapcore.NewTee(infoFileCore)
	} else {
		logCore = zapcore.NewTee(infoConsoleCore)
	}
	logger.infoLogger = zap.New(logCore, zap.AddCaller(), zap.AddCallerSkip(logger.config.CallerSkip), zap.Fields(logger.fields...))

	// err log
	errFileCore := zapcore.NewCore(fileEncoder, errFileWriteSyncer, zap.DebugLevel)
	errConsoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stderr), zap.DebugLevel)

	if !cfg.ConsoleDisabled && !cfg.FileDisabled {
		logCore = zapcore.NewTee(errFileCore, infoFileCore, errConsoleCore)
	} else if !cfg.FileDisabled {
		logCore = zapcore.NewTee(errFileCore, infoFileCore)
	} else {
		logCore = zapcore.NewTee(errConsoleCore)
	}
	logger.errLogger = zap.New(logCore, zap.AddCaller(), zap.AddCallerSkip(logger.config.CallerSkip), zap.Fields(logger.fields...))

	// crash log
	crashFileEncoderConfig := core.DefaultFileEncoderConfig
	crashFileEncoderConfig.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("CRASH")
	}
	crashFileEncoder := zapcore.NewJSONEncoder(crashFileEncoderConfig)
	crashFileCore := zapcore.NewCore(crashFileEncoder, errFileWriteSyncer, zapcore.DebugLevel)

	crashConsoleEncoderConfig := core.DefaultConsoleEncoderConfig
	crashConsoleEncoderConfig.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("\x1b[31m[CRASH]\x1b[0m")
	}
	crashConsoleEncoder := zapcore.NewConsoleEncoder(crashConsoleEncoderConfig)
	crashConsoleCore := zapcore.NewCore(crashConsoleEncoder, zapcore.AddSync(os.Stderr), zap.DebugLevel)

	crashInfoFileCore := zapcore.NewCore(crashFileEncoder, infoFileWriteSyncer, zap.DebugLevel)

	if !cfg.ConsoleDisabled && !cfg.FileDisabled {
		logCore = zapcore.NewTee(crashFileCore, crashInfoFileCore, crashConsoleCore)
	} else if !cfg.FileDisabled {
		logCore = zapcore.NewTee(crashFileCore, crashInfoFileCore)
	} else {
		logCore = zapcore.NewTee(crashInfoFileCore)
	}
	logger.crashLogger = zap.New(logCore, zap.AddCaller(), zap.AddCallerSkip(logger.config.CallerSkip), zap.Fields(logger.fields...))

	return logger, nil
}

func (logger *SysLogger) Info(msg string, fields ...zap.Field) {
	logger.infoLogger.Info(msg, fields...)
}

func (logger *SysLogger) Debug(msg string, fields ...zap.Field) {
	logger.infoLogger.Debug(msg, fields...)
}

func (logger *SysLogger) Warn(msg string, fields ...zap.Field) {
	logger.infoLogger.Warn(msg, fields...)
}

func (logger *SysLogger) Error(msg string, fields ...zap.Field) {
	logger.errLogger.Error(msg, fields...)
}

func (logger *SysLogger) Crash(msg string, fields ...zap.Field) {
	logger.crashLogger.Error(msg, fields...)
}

func (logger *SysLogger) Clone() *SysLogger {
	clone := *logger
	infoClone := *logger.infoLogger
	errClone := *logger.errLogger
	crashClone := *logger.crashLogger
	clone.infoLogger = &infoClone
	clone.errLogger = &errClone
	clone.crashLogger = &crashClone
	return &clone
}

func (logger *SysLogger) With(fields ...zap.Field) {
	logger.infoLogger = logger.infoLogger.With(fields...)
	logger.errLogger = logger.errLogger.With(fields...)
	logger.crashLogger = logger.crashLogger.With(fields...)
}

func (logger *SysLogger) WithOptions(opts ...zap.Option) {
	logger.infoLogger = logger.infoLogger.WithOptions(opts...)
	logger.errLogger = logger.errLogger.WithOptions(opts...)
	logger.crashLogger = logger.crashLogger.WithOptions(opts...)
}

func (logger *SysLogger) WithGID(gid string) {
	if logger.config.GID != "" {
		return
	}
	logger.config.GID = gid
	logger.With(zap.String("gid", gid))
}

func (logger *SysLogger) WithSID(sid int32) {
	if logger.config.SID > 0 {
		return
	}
	logger.config.SID = sid
	logger.With(zap.Int32("sid", sid))
}
