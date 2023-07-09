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

	// People
	r.GET("/people", middlewares.PagedResource, controllers.GetPeople)
	r.GET("/person/:id", controllers.GetPerson)

	r.Run()
}
