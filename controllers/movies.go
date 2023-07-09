package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Preload("Genres").Preload("Director")
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("title LIKE ?", "%"+search+"%").Or("original_title LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &movies, models.Movie{}, filter, searchFilter)
}

func GetMovie(c *gin.Context) {
	var movie models.Movie

	if err := helpers.FirstOrFailWithSlug(c, &movie, c.Param("id"), "Genres", "Director"); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movie,
	})
}

func GetMovieCast(c *gin.Context) {
	var movie models.Movie

	if err := helpers.FirstOrFailWithSlug(c, &movie, c.Param("id")); err != nil {
		return
	}

	var casts []models.Cast

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Where("movie_id = ?", movie.ID).Preload("Person")
	}

	helpers.Paginate(c, &casts, models.Cast{}, filter, nil)
}
