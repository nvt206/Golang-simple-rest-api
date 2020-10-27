package dto

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
	Address string `json:"address"`
	IsAdmin bool `json:"is_admin"`
}
