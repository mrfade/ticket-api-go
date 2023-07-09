package models

import "gorm.io/gorm"

type Theater struct {
	gorm.Model
	Name        string
	Description string
	SeatPlan    string `json:",omitempty"`
	PlaceId     uint
	Place       *Place `gorm:"foreignKey:PlaceId" json:",omitempty"`
}
