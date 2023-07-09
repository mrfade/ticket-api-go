package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func SearchMovies(c *gin.Context) {
	var movies []models.Movie

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Preload("Genres").Preload("Director")
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		db.Where("MATCH(`title`) AGAINST(? IN NATURAL LANGUAGE MODE)", search)
		db.Or("MATCH(`original_title`) AGAINST(? IN NATURAL LANGUAGE MODE)", search)
		db.Or("MATCH(`description`) AGAINST(? IN NATURAL LANGUAGE MODE)", search)
		db.Or("MATCH(`imdb_id`) AGAINST(? IN NATURAL LANGUAGE MODE)", search)
		return db
	}

	helpers.Paginate(c, &movies, models.Movie{}, filter, searchFilter)
}
