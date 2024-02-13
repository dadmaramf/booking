package models

import (
	"github.com/jinzhu/gorm"
)

// Goods model
type Goods struct {
	gorm.Model
	id        int64  `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Price     uint
	Quantity  int32
	Available bool
}
