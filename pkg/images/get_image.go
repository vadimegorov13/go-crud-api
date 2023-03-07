package images

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vadimegorov13/go-crud-api/pkg/common/models"
)

func (h handler) GetImage(c *gin.Context) {
	id := c.Param("id")

	var image models.Image

	if result := h.DB.First(&image, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &image)
}
