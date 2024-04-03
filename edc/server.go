package edc

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func ServerSetup() *fiber.App {
	log.Info("Initializing fiber app")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())

	return app
}
