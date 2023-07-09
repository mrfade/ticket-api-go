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
