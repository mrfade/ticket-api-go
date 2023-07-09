package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/initializers"
	"github.com/mrfade/ticket-api-go/models"
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
