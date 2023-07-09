package models

import (
	"time"

	"gorm.io/gorm"
)

type MovieSession struct {
	gorm.Model
	MovieId   uint
	Movie     *Movie `gorm:"foreignKey:MovieId" json:",omitempty"`
	TheaterId uint
	Theater   *Theater `gorm:"foreignKey:TheaterId" json:",omitempty"`
	ShowTime  time.Time
	Name      string
}
