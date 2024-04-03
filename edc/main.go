package edc

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type EDC struct {
	Config *Configuration
	DB     *gorm.DB
	Cache  *redis.Client
	Http   *fiber.App
}

func (edc *EDC) Serve() {
	log.Info("Starting EDC server")

	if error := edc.Http.Listen(fmt.Sprintf(
		"%s:%d",
		edc.Config.App.Host,
		edc.Config.App.Port,
	)); error != nil {
		log.Fatal(error)
	}
}

var Edc EDC

func Initialize(envfiles ...string) *EDC {
	log.Info("Initializing EDC framework")

	Edc = EDC{
		Config: ConfigSetup(envfiles...),
		DB:     DBSetup(),
		Cache:  CacheSetup(),
		Http:   ServerSetup(),
	}

	return &Edc
}
