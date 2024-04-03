package edc

import (
	"github.com/gofiber/fiber/v2/log"
)

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
