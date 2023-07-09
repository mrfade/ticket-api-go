package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetTheaters(c *gin.Context) {
	var theaters []models.Theater

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("name LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &theaters, models.Theater{}, nil, searchFilter)
}

func GetTheater(c *gin.Context) {
	var theater models.Theater

	if err := helpers.FirstOrFail(c, &theater, c.Param("id")); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": theater,
	})
}

func GetTheaterSeats(c *gin.Context) {
	var seats []models.TheaterSeat

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Where("theater_id = ?", c.Param("id"))
	}

	helpers.Paginate(c, &seats, models.TheaterSeat{}, filter, nil)
}
