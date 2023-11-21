package models

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	Category_ID int    `json:"category_id"`
	User_ID     int    `json:"user_id"`
}
