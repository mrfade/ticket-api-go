package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title            string `gorm:"size:191;index:,class:FULLTEXT"`
	Slug             string `gorm:"size:191;index"`
	OriginalTitle    string `gorm:"size:191;index:,class:FULLTEXT"`
	Description      string `gorm:"index:,class:FULLTEXT"`
	Duration         uint
	PosterPath       string    `gorm:"size:100"`
	BackdropPath     string    `gorm:"size:100"`
	ReleaseDate      time.Time `gorm:"type:date"`
	OriginalLanguage string    `gorm:"size:10"`
	ImdbId           string    `gorm:"size:20;index:,class:FULLTEXT"`
	TmdbId           string    `gorm:"size:20"`
	Status           string    `gorm:"size:20"`
	NowPlaying       bool
	TrailerUrl       string
	Rating           float32
	DirectorId       uint
	Director         *Person `gorm:"foreignKey:DirectorId"`
	Genres           []Genre `gorm:"many2many:movie_genres;"`
	Casts            []Cast  `gorm:"foreignKey:MovieId"`
	Crews            []Crew  `gorm:"foreignKey:MovieId"`
}
