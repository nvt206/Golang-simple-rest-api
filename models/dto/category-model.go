package dto

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" binding:"required,min=3"`
}

