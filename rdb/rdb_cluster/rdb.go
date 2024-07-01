package rdb_cluster

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCluster struct {
	options *RedisClusterOption
	client  *redis.ClusterClient

	ctx    context.Context
	cancel context.CancelFunc
}

func NewRedisCluster() *RedisCluster {
	rs := &RedisCluster{
		options: NewRedisClusterOption(),
	}

	return rs
}

func (r *RedisCluster) Init(ctx context.Context, opts ...any) (err error) {
	r.ctx, r.cancel = context.WithCancel(ctx)
	if len(opts) > 0 {
		for _, opt := range opts {
			opt.(Option)(r.options)
		}
	}
	r.client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:           r.options.addrs,  // 集群地址
		PoolSize:        3000,             // Redis连接池大小
		MaxRetries:      10,               // 最大重试次数
		MinIdleConns:    60,               // 空闲连接数量
		PoolTimeout:     10 * time.Second, // 空闲链接超时时间
		ConnMaxLifetime: 60 * time.Second, // 连接存活时长
		DialTimeout:     15 * time.Second, // 连接建立超时时间，默认5秒。
		ReadTimeout:     7 * time.Second,  // 读超时，默认3秒， -1表示取消读超时
		WriteTimeout:    7 * time.Second,  // 写超时，默认等于读超时
		Password:        r.options.pwd,    // redis 密码
	})

	if err = r.client.Ping(r.ctx).Err(); err != nil {
		return
	}

	return nil

}

func (r *RedisCluster) Get() any {
	return r.client
}

func (r *RedisCluster) GetCtx() context.Context {
	return r.ctx
}
