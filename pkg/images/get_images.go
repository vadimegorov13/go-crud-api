package images

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vadimegorov13/go-crud-api/pkg/common/models"
)

func (h handler) GetImages(c *gin.Context) {
	var images []models.Image

	if result := h.DB.Find(&images); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &images)
}
