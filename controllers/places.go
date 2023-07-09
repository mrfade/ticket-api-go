package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/initializers"
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

func GetPlaceTheaters(c *gin.Context) {
	var theaters []models.Theater

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Where("place_id = ?", c.Param("id"))
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("name LIKE", "%"+search+"%")
	}

	helpers.Paginate(c, &theaters, models.Theater{}, filter, searchFilter)
}

func GetPlaceSessions(c *gin.Context) {
	var sessions []models.MovieSession

	// fetch place's theaters
	var theaters []models.Theater
	if err := initializers.DB.Where("place_id = ?", c.Param("id")).Find(&theaters).Error; err != nil {
		helpers.ErrorJSON(c, http.StatusNotFound, "Not found!")
		return
	}

	if len(theaters) == 0 {
		helpers.ErrorJSON(c, http.StatusNotFound, "Not found!")
		return
	}

	// get theater ids
	var theaterIds []uint
	for _, theater := range theaters {
		theaterIds = append(theaterIds, theater.ID)
	}

	// parse date
	var date time.Time
	if c.Query("date") != "" {
		var err error
		date, err = time.Parse("2006-01-02", c.Query("date"))
		if err != nil {
			helpers.ErrorJSON(c, http.StatusBadRequest, "Invalid date format")
			return
		}
	}

	filter := func(db *gorm.DB) *gorm.DB {
		if !date.IsZero() {
			db.Where("show_time BETWEEN ? AND ?", date.Format(time.DateOnly), date.Add(24*time.Hour).Format(time.DateOnly))
		}
		return db.Where("theater_id IN ?", theaterIds).Preload("Theater")
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("name LIKE", "%"+search+"%")
	}

	helpers.Paginate(c, &sessions, models.MovieSession{}, filter, searchFilter)
}
