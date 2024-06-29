package fnlog

import (
	"time"
)

type Config struct {
	Type string // 日志类型

	Path string // 日志路径

	GID string // 游戏ID
	SID int32  // 服务器 serverID

	LinkEnabled bool // 是否开启打开文件的软连接, 只有linux有效

	ConsoleDisabled bool // 是否关闭终端输出

	FileDisabled bool // 是否关闭文件输出

	MaxAge time.Duration // 最大存放时间, default 1 month

	MaxSize int64 // 最大文件大小, default 300M
}

func (cfg *Config) defalut() *Config {

	if cfg.Path == "" {
		cfg.Path = "./log"
	}

	if cfg.MaxAge <= 0 {
		cfg.MaxAge = 30 * 24 * time.Hour
	}

	if cfg.MaxSize <= 0 {
		cfg.MaxSize = 300 * 1024 * 1024
	}

	return cfg
}
