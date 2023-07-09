package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/initializers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetCities(c *gin.Context) {
	var cities []models.City

	initializers.DB.Find(&cities)

	c.JSON(http.StatusOK, gin.H{
		"data": cities,
	})
}

func GetCity(c *gin.Context) {
	var city models.City

	if err := helpers.FirstOrFail(c, &city, c.Param("id")); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": city,
	})
}

func GetCityPlaces(c *gin.Context) {
	var places []models.Place

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Preload("City").Where("city_id = ?", c.Param("id"))
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("name LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &places, models.Place{}, filter, searchFilter)
}
