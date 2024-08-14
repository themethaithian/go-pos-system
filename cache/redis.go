package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"

	"github.com/themethaithian/go-pos-system/config"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error
}

type cache struct {
	redis *redis.Client
}

func NewRedis() Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Val.RedisAddr,
		Password: config.Val.RedisPassword,
		DB:       config.Val.RedisDB,
	})

	return &cache{
		redis: rdb,
	}
}

func (c *cache) Get(ctx context.Context, key string) (string, error) {
	return c.redis.Get(ctx, key).Result()
}

func (c *cache) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return c.redis.Set(ctx, key, value, exp).Err()
}
