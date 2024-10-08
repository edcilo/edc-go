package edc

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id" redis:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;index" json:"created_at" redis:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index" json:"updated_at" redis:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-" redis:"-"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New()
	return nil
}
