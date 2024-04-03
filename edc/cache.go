package edc

import (
	"context"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

var (
	once          sync.Once
	cacheInstance *redis.Client
)

func CacheSetup() *redis.Client {
	log.Info("Setting up Cache connection...")

	engine := Config.Cache.Engine
	dsn := CacheDSN{
		Host:     Config.Cache.Host,
		Port:     Config.Cache.Port,
		User:     Config.Cache.User,
		Password: Config.Cache.Password,
		Database: Config.Cache.Database,
	}

	switch engine {
	case Redis:
		once.Do(func() {
			cacheInstance = RedisConn(dsn)
		})

		ctx := context.Background()
		pong := cacheInstance.Ping(ctx)
		if pong.Err() != nil {
			msg := fmt.Sprintf("Failed to connect to Cache: %s", pong.Err())
			panic(msg)
		}
	default:
		msg := fmt.Sprintf(
			"Invalid Cache Engine. Valid options are: %s. You provided: %s",
			Redis, engine)
		panic(msg)
	}

	log.Info("Cache connection established.")
	return cacheInstance
}

func RedisConn(params CacheDSN) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", params.Host, params.Port),
		DB:       params.Database,
		Username: params.User,
		Password: params.Password,
	})
	return client
}
