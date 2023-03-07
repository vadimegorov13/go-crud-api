package models

import (
	"gorm.io/gorm"
)

/*
gorm.model will add properties such as ID, CreatedAt, UpdatedAt and DeletedAt
*/

type User struct {
	gorm.Model
	Username  string `json:"username"`
	ImagesNum int    `json:"images_num"`
}
