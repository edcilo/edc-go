package edc

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ErrorResponse(err error, c *fiber.Ctx) error {
	status := fiber.StatusInternalServerError
	response := fiber.Map{
		"message": err.Error(),
	}

	if err == gorm.ErrRecordNotFound {
		status = fiber.StatusNotFound
	}

	return c.Status(status).JSON(response)
}
