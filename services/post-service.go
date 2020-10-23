package services

import (
	"demo/common"
	"demo/models"
	"github.com/jinzhu/gorm"
	"sync"
)

var postInit sync.Once
var postInstance *postService

type PostService interface {

	Post(post *models.Post) *models.Post
	GetByUserAndCategory(userid,categoryid uint) []models.Post
	FindListPost(condition interface{}) []models.Post

}
type postService struct {
	DB *gorm.DB
}

func (p postService) FindListPost(condition interface{}) []models.Post {
	var posts []models.Post
	if err := p.DB.Debug().Where(condition).Find(&posts).Error;err!=nil {
		return nil
	}
	return posts
}

func (p postService) GetByUserAndCategory(userid, categoryid uint) []models.Post {

	var posts []models.Post
	if err := p.DB.Debug().Where("user_id=? and category_id=?",userid,categoryid).Find(&posts).Error;err!=nil {
		return nil
	}
	return posts
}


func (p postService) Post(post *models.Post) *models.Post {
	if err := p.DB.Create(&post).Error;err!=nil{
		return nil
	}
	return post
}

func NewPostService() PostService {

	if postInstance == nil{
		postInit.Do(
			func() {
				postInstance = &postService{DB: common.ConnectData()}
			})
	}
	return postInstance
}