package models

import "gorm.io/gorm"

type MovieSessionSeat struct {
	gorm.Model
	SessionId uint
	Session   *MovieSession `gorm:"foreignKey:SessionId" json:",omitempty"`
	SeatId    uint
	Seat      *TheaterSeat `gorm:"foreignKey:SeatId" json:",omitempty"`
	UserId    uint
	User      *User `gorm:"foreignKey:UserId" json:",omitempty"`
}
