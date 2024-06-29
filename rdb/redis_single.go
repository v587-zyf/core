package rdb

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisSingle struct {
	client *redis.Client
	Ctx    context.Context
}

func NewRedisSingle(addr, password string) (*RedisSingle, error) {
	rs := &RedisSingle{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
			PoolSize: 300,
		}),
		Ctx: context.Background(),
	}

	if err := rs.client.Ping(rs.Ctx).Err(); err != nil {
		return nil, err
	}
	return rs, nil
}

func (r *RedisSingle) Get() *redis.Client {
	return r.client
}

func (r *RedisSingle) GetCtx() context.Context {
	return r.Ctx
}
