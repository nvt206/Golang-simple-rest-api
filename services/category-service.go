package services

import (
	"demo/common"
	"github.com/jinzhu/gorm"
	"sync"
)

var categoryInstance *categoryService
var categoryOnce sync.Once

type CategoryService interface {
	//GetAll() []dto.Category
	//GetById(id uint) *dto.Category
	//Post(category *dto.Category) *dto.Category
	//Delete(ctx *gin.Context,id uint) error
}
//
type categoryService struct {
	DB *gorm.DB

}
//
//func (c categoryService) Delete(ctx *gin.Context, id uint) error {
//	return c.DB.Where("ID=?",id).Delete(&dto.Category{}).Error
//
//}
//
//func (c categoryService) Post(category *dto.Category) *dto.Category {
//	if err := c.DB.Create(&category).Error;err!=nil{
//		return nil
//	}
//	return category
//}
//
//func (c categoryService) GetAll() []dto.Category {
//	var categories []dto.Category
//	c.DB.Find(&categories)
//	return categories
//}
//
//func (c categoryService) GetById(id uint) *dto.Category {
//	var category dto.Category
//	if err := c.DB.Where("id=?",id).Find(&category).Error;err !=nil{
//		return nil
//	}
//	return &category
//}

func NewCategoryService() CategoryService {

	if categoryInstance == nil{
		categoryOnce.Do(func() {
			categoryInstance = &categoryService{DB: common.ConnectData()}
		})
	}
	return categoryInstance
}
