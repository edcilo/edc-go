package helpers

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB, models []interface{}) {
	log.Info("🤖 Executing migrations...")
	for _, model := range models {
		db.AutoMigrate(model)
	}
	log.Info("🏁 Migrations executed successfully.")
}
