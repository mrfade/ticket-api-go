package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/initializers"
	"github.com/mrfade/ticket-api-go/models"
	"gorm.io/gorm"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Preload("Genres").Preload("Director")
	}

	searchFilter := func(db *gorm.DB, search string) *gorm.DB {
		return db.Where("title LIKE ?", "%"+search+"%").Or("original_title LIKE ?", "%"+search+"%")
	}

	helpers.Paginate(c, &movies, models.Movie{}, filter, searchFilter)
}

func GetMovie(c *gin.Context) {
	var movie models.Movie

	if err := helpers.FirstOrFailWithSlug(c, &movie, c.Param("id"), "Genres", "Director"); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movie,
	})
}

func GetMovieCast(c *gin.Context) {
	var movie models.Movie

	if err := helpers.FirstOrFailWithSlug(c, &movie, c.Param("id")); err != nil {
		return
	}

	var casts []models.Cast

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Where("movie_id = ?", movie.ID).Preload("Person")
	}

	helpers.Paginate(c, &casts, models.Cast{}, filter, nil)
}

func GetMovieCrew(c *gin.Context) {
	var movie models.Movie

	if err := helpers.FirstOrFailWithSlug(c, &movie, c.Param("id")); err != nil {
		return
	}

	var crews []models.Crew

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Where("movie_id = ?", movie.ID).Preload("Person")
	}

	helpers.Paginate(c, &crews, models.Crew{}, filter, nil)
}

func GetMovieSessions(c *gin.Context) {
	var movie models.Movie

	if err := helpers.FirstOrFailWithSlug(c, &movie, c.Param("id")); err != nil {
		return
	}

	var date time.Time
	var cityId uint

	// parse date
	qdate := c.Query("date")
	if qdate != "" {
		var err error
		if date, err = time.Parse("2006-01-02", qdate); err != nil {
			helpers.ErrorJSON(c, http.StatusBadRequest, "Invalid date format")
			return
		}
	}

	// parse city
	qcity := c.Query("city")
	if qcity != "" {
		if city, err := strconv.Atoi(qcity); err == nil {
			cityId = uint(city)
		}
	}

	if cityId == 0 || date.IsZero() {
		helpers.ErrorJSON(c, http.StatusBadRequest, "Invalid query")
		return
	}

	var places []models.Place

	filter := func(db *gorm.DB) *gorm.DB {
		if cityId != 0 {
			db.Where("city_id = ?", cityId)
		}

		return db.Preload("Theaters", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "description", "place_id")
		}).Preload("Theaters.MovieSessions", func(db *gorm.DB) *gorm.DB {
			return db.Where("movie_id = ?", movie.ID).Where("show_time BETWEEN ? AND ?", date.Format(time.DateOnly), date.Add(24*time.Hour).Format(time.DateOnly))
		})
	}

	helpers.Paginate(c, &places, models.Place{}, filter, nil)
}

func GetSimilarMovies(c *gin.Context) {
	var movies []models.Movie

	initializers.DB.Preload("Genres").Preload("Director").Order("rand()").Limit(6).Find(&movies)

	c.JSON(http.StatusOK, gin.H{
		"data": movies,
	})
}
