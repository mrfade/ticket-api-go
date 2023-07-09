package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/controllers"
	"github.com/mrfade/ticket-api-go/initializers"
	"github.com/mrfade/ticket-api-go/middlewares"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Ticket API",
		})
	})

	// Casts
	r.GET("/casts", middlewares.PagedResource, controllers.GetCasts)
	r.GET("/cast/:id", controllers.GetCast)

	// Cities
	r.GET("/cities", controllers.GetCities)
	r.GET("/city/:id", controllers.GetCity)
	r.GET("/city/:id/places", middlewares.PagedResource, controllers.GetCityPlaces)

	// Crews
	r.GET("/crews", middlewares.PagedResource, controllers.GetCrews)
	r.GET("/crew/:id", controllers.GetCrew)

	// Genres
	r.GET("/genres", middlewares.PagedResource, controllers.GetGenres)
	r.GET("/genre/:id", controllers.GetGenre)

	// Movies
	r.GET("/movies", middlewares.PagedResource, controllers.GetMovies)
	r.GET("/movie/:id", controllers.GetMovie)
	r.GET("/movie/:id/cast", middlewares.PagedResource, controllers.GetMovieCast)
	r.GET("/movie/:id/crew", middlewares.PagedResource, controllers.GetMovieCrew)
	r.GET("/movie/:id/sessions", middlewares.PagedResource, controllers.GetMovieSessions)
	r.GET("/movie/:id/similar_movies", controllers.GetSimilarMovies)

	// People
	r.GET("/people", middlewares.PagedResource, controllers.GetPeople)
	r.GET("/person/:id", controllers.GetPerson)
	r.GET("/person/:id/movies", middlewares.PagedResource, controllers.GetPersonMovies)

	// Places
	r.GET("/places", middlewares.PagedResource, controllers.GetPlaces)
	r.GET("/place/:id", controllers.GetPlace)
	r.GET("/place/:id/sessions", middlewares.PagedResource, controllers.GetPlaceSessions)
	r.GET("/place/:id/theaters", middlewares.PagedResource, controllers.GetPlaceTheaters)

	// Search
	r.GET("/search/movie", middlewares.PagedResource, controllers.SearchMovies)

	// Sessions
	r.GET("/sessions", middlewares.PagedResource, controllers.GetSessions)
	r.GET("/session/:id", controllers.GetSession)

	// Theaters
	r.GET("/theaters", middlewares.PagedResource, controllers.GetTheaters)
	r.GET("/theater/:id", controllers.GetTheater)
	r.GET("/theater/:id/seats", middlewares.PagedResource, controllers.GetTheaterSeats)
	r.GET("/theater/:id/sessions", middlewares.PagedResource, controllers.GetTheaterSessions)

	r.Run()
}
