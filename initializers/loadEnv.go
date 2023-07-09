package initializers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil && gin.Mode() != "release" {
		log.Fatal("Error loading .env file")
	}
}
