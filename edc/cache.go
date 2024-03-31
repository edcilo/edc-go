package edc

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

var (
	once          sync.Once
	cacheInstance *redis.Client
)

func CacheSetup(params CacheSetupArgs) *redis.Client {
	log.Info("Setting up Cache connection...")

	switch params.Engine {
	case Redis:
		once.Do(func() {
			cacheInstance = RedisConn(params.DSN)
		})
	default:
		msg := fmt.Sprintf(
			"Invalid Cache Engine. Valid options are: %s. You provided: %s",
			Redis, params.Engine)
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
