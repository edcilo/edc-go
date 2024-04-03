package edc

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type EDC struct {
	Config *Configuration
	DB     *gorm.DB
	Cache  *redis.Client
}

var Edc EDC

func Initialize(envfiles ...string) *EDC {
	log.Info("Initializing EDC framework")

	Edc = EDC{
		Config: ConfigSetup(envfiles...),
		DB:     DBSetup(),
		Cache:  CacheSetup(),
	}

	return &Edc
}
