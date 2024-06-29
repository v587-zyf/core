package log

import (
	"core/log/syslog"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	GID string

	Fields []zapcore.Field

	LinkEnabled bool

	ConsoleEnabled bool
	FileEnabled    bool

	CallerSkip int

	SysLevel zapcore.Level

	SysPath    string
	AccessPath string
	StatPath   string

	SysMaxAge    time.Duration
	AccessMaxAge time.Duration
	StatMaxAge   time.Duration

	MaxSize int64

	config *Config

	sysLogger *syslog.SysLogger
}

func New(cfg *Config) (*Logger, error) {
	if cfg.CallerSkip == 0 {
		cfg.CallerSkip = defConfig.CallerSkip
	}
	if cfg.SysLevel == 0 {
		cfg.SysLevel = defConfig.SysLevel
	}
	if cfg.Path == "" {
		cfg.Path = defConfig.Path
	}

	if cfg.SysMaxAge == 0 {
		cfg.SysMaxAge = defConfig.SysMaxAge
	}

	if cfg.MaxSize == 0 {
		cfg.MaxSize = defConfig.MaxSize
	}

	logger := &Logger{
		config: cfg,
	}

	var err error
	logger.sysLogger, err = syslog.New(&syslog.Config{
		Path:            cfg.Path,
		GID:             cfg.GID,
		ConsoleDisabled: cfg.ConsoleDisabled,
		FileDisabled:    cfg.FileDisabled,
	})
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func (logger *Logger) Info(msg string, fields ...zap.Field) {
	logger.sysLogger.Info(msg, fields...)
}

func (logger *Logger) Debug(msg string, fields ...zap.Field) {
	logger.sysLogger.Debug(msg, fields...)
}

func (logger *Logger) Warn(msg string, fields ...zap.Field) {
	logger.sysLogger.Warn(msg, fields...)
}

func (logger *Logger) Error(msg string, fields ...zap.Field) {
	logger.sysLogger.Error(msg, fields...)
}

func (logger *Logger) Crash(msg string, fields ...zap.Field) {
	logger.sysLogger.Crash(msg, fields...)
}

func (logger *Logger) Clone() *Logger {
	clone := *logger
	configClone := *logger.config
	clone.config = &configClone

	clone.sysLogger = logger.sysLogger.Clone()

	return &clone
}

func (logger *Logger) With(fields ...zap.Field) {
	logger.sysLogger.With(fields...)
}

func (logger *Logger) WithOptions(opts ...zap.Option) {
	logger.sysLogger.WithOptions(opts...)
}

func (logger *Logger) WithGID(gid string) {
	logger.sysLogger.WithGID(gid)
}

func (logger *Logger) WithSID(sid int32) {
	logger.sysLogger.WithSID(sid)
}
