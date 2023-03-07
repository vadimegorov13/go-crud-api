package images

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
The image handlers/route will be based on Pointer receivers, for
that, it is defined as struct. This struct will receive the database
information later, so whenever image handler/route is called gives
the access to GORM.
*/

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/images")

	routes.POST("/", h.AddImage)
	routes.GET("/", h.GetImages)
	routes.GET("/:id", h.GetImage)
	routes.PATCH("/:id", h.UpdateImage)
	routes.DELETE("/:id", h.DeleteImage)
}
