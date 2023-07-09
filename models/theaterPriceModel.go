package models

import "gorm.io/gorm"

type TheaterPrice struct {
	gorm.Model
	TheaterId uint
	Theater   *Theater `gorm:"foreignKey:TheaterId"`
	Price     uint
	Type      uint
}
