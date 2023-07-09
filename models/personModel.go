package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	TmdbId             uint   `gorm:"unique"`
	ImdbId             string `gorm:"unique;size:20"`
	Name               string
	Slug               string `gorm:"unique;index"`
	Biography          string
	Birthday           string
	PlaceOfBirth       string
	Gender             uint
	KnownForDepartment string
	ProfilePath        string
}
