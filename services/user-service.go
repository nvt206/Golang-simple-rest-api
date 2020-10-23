package services

import (
	"demo/common"
	"demo/models"
	"github.com/jinzhu/gorm"
	"sync"
)
var userInit sync.Once
var userInstance *userService



type UserService interface {

	FindOne(condition interface{}) *models.User
	FindById(id uint) *models.User
	GetAll() []models.User
	Register(user *models.User) *models.User
	Update(user *models.User) *models.User


}

type userService struct {
	 DB *gorm.DB
}

func (u userService) FindById(id uint) *models.User {
	var user models.User
	if err := u.DB.Where("id=?",id).Find(&user).Error;err !=nil{
		return nil
	}
	return &user
}

func (u userService) Update(user *models.User) *models.User {

	if err := u.DB.Save(&user).Error;err!=nil{
		return nil
	}
	return user
}

func (u userService) Register(user *models.User) *models.User {
	if err := u.DB.Create(&user).Error;err!=nil{
		return nil
	}
	return user
}

func (u userService) GetAll() []models.User {
	var users []models.User
	u.DB.Find(&users)
	return users
}

func (u userService) FindOne(condition interface{}) *models.User {
	var user models.User
	if err := u.DB.Debug().Where(condition).Find(&user).Error;err !=nil{
		return nil
	}
	return &user
}

func NewUserService() UserService {
	if userInstance == nil{
		userInit.Do(
			func() {
				userInstance = &userService{DB: common.ConnectData()}
			},
		)
	}
	return userInstance
}