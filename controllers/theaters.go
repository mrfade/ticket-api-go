package controllers

import (
	"net/http"
	"time"

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

func GetTheaterSessions(c *gin.Context) {
	var sessions []models.MovieSession

	var date time.Time
	if c.Query("date") != "" {
		date, _ = time.Parse("2006-01-02", c.Query("date"))
	}

	filter := func(db *gorm.DB) *gorm.DB {
		if !date.IsZero() {
			db.Where("show_time BETWEEN ? AND ?", date.Format(time.DateOnly), date.Add(24*time.Hour).Format(time.DateOnly))
		}
		return db.Where("theater_id = ?", c.Param("id"))
	}

	helpers.Paginate(c, &sessions, models.MovieSession{}, filter, nil)
}
