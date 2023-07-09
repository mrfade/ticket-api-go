package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserId     uint
	User       *User `gorm:"foreignKey:UserId"`
	SessionId  uint
	Session    *MovieSession `gorm:"foreignKey:SessionId"`
	Seats      []TheaterSeat `gorm:"many2many:movie_session_seats;"`
	TotalPrice uint
}
