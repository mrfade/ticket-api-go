package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetPlaces(c *gin.Context) {
	var places []models.Place

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Preload("City")
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("name LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &places, models.Place{}, filter, searchFilter)
}

func GetPlace(c *gin.Context) {
	var place models.Place

	if err := helpers.FirstOrFail(c, &place, c.Param("id"), "City"); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": place,
	})
}
