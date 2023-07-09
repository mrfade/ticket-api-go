package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetCasts(c *gin.Context) {
	var casts []models.Cast

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Preload("Person")
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("role LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &casts, models.Cast{}, filter, searchFilter)
}

func GetCast(c *gin.Context) {
	var cast models.Cast

	if err := helpers.FirstOrFail(c, &cast, c.Param("id"), "Person"); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": cast,
	})
}
