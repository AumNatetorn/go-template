package cache

import (
	"context"
	"fmt"
	"go-template/configs"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type cache struct {
	client *redis.Client
}

func NewCacheClient(conf configs.Redis, secrets configs.Secrets) (*cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Addr, conf.Port),
		Password: secrets.RedisPassword,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, errors.Wrap(err, "new cache client")
	}

	return &cache{client: client}, nil
}

func (c *cache) Close() error {
	return c.client.Close()
}

func (client *cache) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.client.Ping(ctx).Err()
}
