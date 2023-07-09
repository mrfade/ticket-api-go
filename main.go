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

	// Crews
	r.GET("/crews", middlewares.PagedResource, controllers.GetCrews)
	r.GET("/crew/:id", controllers.GetCrew)

	// Genres
	r.GET("/genres", middlewares.PagedResource, controllers.GetGenres)
	r.GET("/genre/:id", controllers.GetGenre)

	// Movies
	r.GET("/movies", middlewares.PagedResource, controllers.GetMovies)
	r.GET("/movie/:id", controllers.GetMovie)

	// People
	r.GET("/people", middlewares.PagedResource, controllers.GetPeople)
	r.GET("/person/:id", controllers.GetPerson)

	// Places
	r.GET("/places", middlewares.PagedResource, controllers.GetPlaces)
	r.GET("/place/:id", controllers.GetPlace)

	r.Run()
}
