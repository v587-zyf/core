package rdb_cluster

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var defRedis *RedisCluster

func InitCluster(ctx context.Context, opts ...any) (err error) {
	defRedis = NewRedisCluster()
	if err = defRedis.Init(ctx, opts...); err != nil {
		return err
	}

	return nil
}

func Get() *redis.ClusterClient {
	return defRedis.Get().(*redis.ClusterClient)
}

func GetCtx() context.Context {
	return defRedis.GetCtx()
}
