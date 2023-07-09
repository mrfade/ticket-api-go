package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetCrews(c *gin.Context) {
	var crews []models.Crew

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Preload("Person")
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("job LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &crews, models.Crew{}, filter, searchFilter)
}

func GetCrew(c *gin.Context) {
	var crew models.Crew

	if err := helpers.FirstOrFail(c, &crew, c.Param("id"), "Person"); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": crew,
	})
}
