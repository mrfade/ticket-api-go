package models

import "gorm.io/gorm"

type TheaterSeat struct {
	gorm.Model
	TheaterId uint
	Theater   *Theater `gorm:"foreignKey:TheaterId" json:",omitempty"`
	Row       string
	Number    uint
	Name      string
}
