package rdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)

var defRedis *RedisSingle
var defRedisMu sync.Mutex

func Init(cfg *ConfigSingle) (err error) {
	defRedisMu.Lock()
	defer defRedisMu.Unlock()

	defRedis, err = NewRedisSingle(cfg.Addr, cfg.Pwd)
	if err != nil {
		return err
	}

	return nil
}

func Get() *redis.Client {
	return defRedis.Get()
}

func GetCtx() context.Context {
	return defRedis.Ctx
}
