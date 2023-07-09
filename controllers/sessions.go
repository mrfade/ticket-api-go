package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetSessions(c *gin.Context) {
	var sessions []models.MovieSession

	var date time.Time
	if c.Query("date") != "" {
		date, _ = time.Parse("2006-01-02", c.Query("date"))
	}

	filter := func(db *gorm.DB) *gorm.DB {
		if !date.IsZero() {
			db.Where("show_time BETWEEN ? AND ?", date.Format(time.DateOnly), date.Add(24*time.Hour).Format(time.DateOnly))
		}
		return db
	}

	helpers.Paginate(c, &sessions, models.MovieSession{}, filter, nil)
}

func GetSession(c *gin.Context) {
	var session models.MovieSession

	if err := helpers.FirstOrFail(c, &session, c.Param("id")); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": session,
	})
}
