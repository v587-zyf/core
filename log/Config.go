package log

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Config struct {
	GID string

	Fields []zapcore.Field

	LinkEnabled bool

	ConsoleDisabled bool
	FileDisabled    bool

	CallerSkip int

	SysLevel zapcore.Level

	Path string

	SysMaxAge time.Duration

	MaxSize int64
}
