package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vadimegorov13/go-crud-api/pkg/common/db"
	"github.com/vadimegorov13/go-crud-api/pkg/images"
)

/*
Initialize viper to handle env variables
*/

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	// add env variables as needed
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	images.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
	})

	router.Run(port)
}
