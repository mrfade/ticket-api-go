package models

import "gorm.io/gorm"

type Crew struct {
	gorm.Model
	PersonId   uint
	Person     *Person `gorm:"foreignKey:PersonId"`
	MovieId    uint
	Department string
	Job        string
}
