package model

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	UserId uint `json:"user_id"`
}

