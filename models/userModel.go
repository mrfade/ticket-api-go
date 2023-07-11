package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string  `gorm:"unique"`
	Password string  `json:"-"`
	Status   string  `gorm:"default:active"`
	Roles    *[]Role `gorm:"many2many:user_roles;"`
}
