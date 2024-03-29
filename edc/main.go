package edc

import (
	"gorm.io/gorm"
)

type EDC struct {
	DB *gorm.DB
}

func NewEDC(args NewEDCArgs) *EDC {
	return &EDC{
		DB: DBSetup(args.DB),
	}
}
