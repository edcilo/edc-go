package edc

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" redis:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at" redis:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at" redis:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-" redis:"-"`
}
