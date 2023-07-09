package models

import "gorm.io/gorm"

type Cast struct {
	gorm.Model
	PersonId uint
	Person   *Person `gorm:"foreignKey:PersonId"`
	MovieId  uint
	Role     string
}
