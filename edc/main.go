package edc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type EDC struct {
	DB    *gorm.DB
	Cache *redis.Client
}

var edc = EDC{}

func Initialize(args NewEDCArgs) EDC {
	edc = EDC{
		DB:    DBSetup(args.DB),
		Cache: CacheSetup(args.Cache),
	}
	return edc
}
