package images

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vadimegorov13/go-crud-api/pkg/common/models"
)

type AddImageRequestBody struct {
	Title      string `json:"title"`
	Resolution string `json:"resolution"`
	AIModel    string `json:"ai_model"`
	Prompt     string `json:"prompt"`
	URL        string `json:"url"`
}

func (h handler) AddImage(c *gin.Context) {
	body := AddImageRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var image models.Image

	image.Title = body.Title
	image.Resolution = body.Resolution
	image.AIModel = body.AIModel
	image.Prompt = body.Prompt
	image.URL = body.URL

	if result := h.DB.Create(&image); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &image)
}
