package helpers

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/initializers"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context, data interface{}, model interface{}, filter func(db *gorm.DB) *gorm.DB, searchFilter func(db *gorm.DB, search string) *gorm.DB) {
	pageNumber := c.GetInt("pageNumber")
	pageSize := c.GetInt("pageSize")
	search := c.Query("q")

	// check if pageNumber and pageSize are valid
	if pageNumber < 1 {
		pageNumber = 1
	}

	if pageSize < 10 {
		pageSize = 10
	}

	var dataQuery = initializers.DB.Model(&model)
	if filter != nil {
		dataQuery = filter(dataQuery)
	}
	if search != "" && searchFilter != nil {
		dataQuery = searchFilter(dataQuery, search)
	}

	dataQuery.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(data)

	var totalRecords int64
	var totalQuery = initializers.DB.Model(&model)
	if filter != nil {
		totalQuery = filter(totalQuery)
	}
	if search != "" && searchFilter != nil {
		totalQuery = searchFilter(totalQuery, search)
	}
	totalQuery.Count(&totalRecords)

	totalPages := math.Ceil(float64(totalRecords) / float64(pageSize))

	c.JSON(http.StatusOK, gin.H{
		"pageNumber":   pageNumber,
		"pageSize":     pageSize,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
		"hasNext":      pageNumber < int(totalPages),
		"hasPrev":      pageNumber > 1,
		"data":         data,
	})
}
