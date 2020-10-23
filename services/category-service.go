package services

import (
	"demo/common"
	"demo/models"
	"github.com/jinzhu/gorm"
)

type CategoryService interface {
	GetAll() []models.Category
	GetById(id uint) *models.Category
	Post(category *models.Category) *models.Category
}

type categoryService struct {
	DB *gorm.DB
}

func (c categoryService) Post(category *models.Category) *models.Category {
	if err := c.DB.Create(&category).Error;err!=nil{
		return nil
	}
	return category
}

func (c categoryService) GetAll() []models.Category {
	var categories []models.Category
	c.DB.Find(&categories)
	return categories
}

func (c categoryService) GetById(id uint) *models.Category {
	var category models.Category
	if err := c.DB.Where("id=?",id).Find(&category).Error;err !=nil{
		return nil
	}
	return &category
}

func NewCategoryService() CategoryService {
	return &categoryService{DB: common.ConnectData()}
}
