package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/ticket-api-go/initializers"
)

func FirstOrFail(c *gin.Context, model interface{}, id string, preload ...string) error {
	var query = initializers.DB

	for _, v := range preload {
		query = query.Preload(v)
	}

	query.First(model, id)

	if query.Error != nil {
		ErrorJSON(c, http.StatusNotFound, "Not found!")
		return query.Error
	}

	return nil
}

func FirstOrFailWithSlug(c *gin.Context, model interface{}, id string, preload ...string) error {
	var query = initializers.DB

	for _, v := range preload {
		query = query.Preload(v)
	}

	query.Where("id = ?", id).Or("slug = ?", id).First(model)

	if query.Error != nil {
		ErrorJSON(c, http.StatusNotFound, "Not found!")
		return query.Error
	}

	return nil
}
