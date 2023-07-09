package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	TmdbId uint   `gorm:"unique"`
	Name   string `gorm:"unique"`
}
