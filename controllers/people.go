package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetPeople(c *gin.Context) {
	var people []models.Person

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("name LIKE ?", "%"+search+"%").Or("imdb_id LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &people, models.Person{}, nil, searchFilter)
}

func GetPerson(c *gin.Context) {
	var person models.Person

	if err := helpers.FirstOrFailWithSlug(c, &person, c.Param("id")); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": person,
	})
}

func GetPersonMovies(c *gin.Context) {
	var person models.Person

	if err := helpers.FirstOrFailWithSlug(c, &person, c.Param("id")); err != nil {
		return
	}

	var movies []models.Movie

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Joins("JOIN casts ON casts.movie_id = movies.id").Where("casts.person_id = ?", person.ID)
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("title LIKE ?", "%"+search+"%").Or("original_title LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &movies, models.Movie{}, filter, searchFilter)
}
