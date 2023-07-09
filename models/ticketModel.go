package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserId     uint
	User       *User `gorm:"foreignKey:UserId"`
	SessionId  uint
	Session    *MovieSession `gorm:"foreignKey:SessionId"`
	TotalPrice uint
}
