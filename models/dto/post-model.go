package dto

import (
	"demo/models"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Content string            `json:"content" binding:"required"`
	CategoryId uint           `json:"category_id" binding:"required"`
	UserId uint               `json:"user_id"`
	Category *Category        `json:"category"`
	UserInfo *models.UserInfo `json:"user_info"`
}

