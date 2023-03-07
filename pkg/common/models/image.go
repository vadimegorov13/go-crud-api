package models

import (
	"gorm.io/gorm"
)

/*
gorm.model will add properties such as ID, CreatedAt, UpdatedAt and DeletedAt
*/

type Image struct {
	gorm.Model
	Title      string `json:"title"`
	Resolution string `json:"resolution"`
	AIModel    string `json:"ai_model"`
	Prompt     string `json:"prompt"`
	URL        string `json:"url"`
	Author     *User  `json:"author"`
}
