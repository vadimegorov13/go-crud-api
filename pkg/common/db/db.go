package db

import (
	"log"

	"github.com/vadimegorov13/go-crud-api/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
GORM configuration and auto migration of the schema we just created.
AutoMigrate function will create the images and users table in the
database as soon as application is ran.
*/

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		// this will be printed in the terminal, confirming the connection to the database
		log.Println("The database is connected")
	}

	db.AutoMigrate(&models.Image{})

	return db
}
