package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PagedResource(c *gin.Context) {
	// get the page number and size
	pageNumber := 1
	pageSize := 10

	if i, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1")); err == nil {
		pageNumber = i
	}

	if i, err := strconv.Atoi(c.DefaultQuery("pageSize", "10")); err == nil {
		pageSize = i
	}

	if pageNumber < 1 {
		pageNumber = 1
	}

	if pageSize < 10 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	// set the page number and size
	c.Set("pageNumber", pageNumber)
	c.Set("pageSize", pageSize)

	// continue
	c.Next()
}
