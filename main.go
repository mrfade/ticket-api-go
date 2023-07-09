package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/controllers"
	"github.com/mrfade/ticket-api-go/initializers"
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

	// Cities
	r.GET("/cities", controllers.GetCities)
	r.GET("/city/:id", controllers.GetCity)

	r.Run()
}
