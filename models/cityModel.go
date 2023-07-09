package models

type City struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
