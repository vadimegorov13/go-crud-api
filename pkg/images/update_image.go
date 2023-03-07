package images

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vadimegorov13/go-crud-api/pkg/common/models"
)

type UpdateImageRequestBody struct {
	Title      string `json:"title"`
	Resolution string `json:"resolution"`
	AIModel    string `json:"ai_model"`
	Prompt     string `json:"prompt"`
	URL        string `json:"url"`
}

func (h handler) UpdateImage(c *gin.Context) {
	id := c.Param("id")
	body := AddImageRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var image models.Image

	if result := h.DB.First(&image, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	image.Title = body.Title
	image.Resolution = body.Resolution
	image.AIModel = body.AIModel
	image.Prompt = body.Prompt
	image.URL = body.URL

	h.DB.Save(&image)

	c.JSON(http.StatusCreated, &image)
}
