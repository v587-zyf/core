package syslog

import (
	"time"
)

type Config struct {
	Path string

	GID string
	SID int32

	LinkEnabled bool

	CallerSkip int

	ConsoleDisabled bool
	FileDisabled    bool

	MaxAge  time.Duration // default 1 month
	MaxSize int64         // default 300M
}

func (cfg *Config) get() *Config {

	if cfg.Path == "" {
		cfg.Path = "./log"
	}

	if cfg.MaxAge <= 0 {
		cfg.MaxAge = 30 * 24 * time.Hour
	}

	if cfg.MaxSize <= 0 {
		cfg.MaxSize = 300 * 1024 * 1024
	}

	if cfg.CallerSkip <= 0 {
		cfg.CallerSkip = 3
	}

	return cfg
}
