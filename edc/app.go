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

func (edc *EDC) DBMigrate(models []interface{}) {
	DBMigrate(edc.DB, models)
}

func (edc *EDC) DBRunSeeders(seeders []DBSeederFunc) {
	DBRunSeeders(edc.DB, seeders)
}
