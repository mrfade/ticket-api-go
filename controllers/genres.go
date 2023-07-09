package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetGenres(c *gin.Context) {
	var genres []models.Genre

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("name LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &genres, models.Genre{}, nil, searchFilter)
}

func GetGenre(c *gin.Context) {
	var genre models.Genre

	if err := helpers.FirstOrFail(c, &genre, c.Param("id")); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": genre,
	})
}
