package models

import "github.com/jinzhu/gorm"

type UserInfo struct {
	gorm.Model
	Email string `json:"email" binding:"required,email"`
	Address string `json:"address"`

}