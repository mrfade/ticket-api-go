package models

import "gorm.io/gorm"

type MovieGenre struct {
	gorm.Model
	MovieId uint
	Movie   *Movie `gorm:"foreignKey:MovieId"`
	GenreId uint
	Genre   *Genre `gorm:"foreignKey:GenreId"`
}
