package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	// get the user from the context
	user, _ := c.Get("user")

	// response
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
