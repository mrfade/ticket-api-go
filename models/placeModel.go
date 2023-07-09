package models

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Name        string
	Description string
	CityId      uint
	City        *City `gorm:"foreignKey:CityId" json:",omitempty"`
	Address     string
	Latitude    float64
	Longitude   float64
}
