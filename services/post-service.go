package services

import (
	"demo/common"
	"demo/models/dto"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"sync"
)

var postOnce sync.Once
var postInstance *postService

type PostService interface {

	//Post(ctx *gin.Context,post *dto.Post) (*dto.Post,error)
	//FindListPost(condition interface{}) ([]dto.Post,error)
	MapUserAndCategory(post *dto.Post)

}
type postService struct {
	DB *gorm.DB
}

func (p postService) MapUserAndCategory(post *dto.Post) {
	post.Category =p.DB.Where("ID=?",post.CategoryId).Find(&dto.Category{}).Value.(*dto.Category)
	mapstructure.Decode(p.DB.Where("ID=?",post.UserId).Find(&dto.User{}).Value.(*dto.User),&post.UserInfo)
}

//func (p postService) FindListPost(condition interface{}) ([]dto.Post,error) {
//	var posts []dto.Post
//	if err := p.DB.Debug().Where(condition).Find(&posts).Error;err!=nil {
//		return nil,err
//	}
//	for i,_ :=range posts {
//		posts[i].Category =p.DB.Debug().Where("ID=?",posts[i].CategoryId).Find(&dto.Category{}).Value.(*dto.Category)
//		mapstructure.Decode(p.DB.Debug().Where("ID=?",posts[i].UserId).Find(&dto.User{}).Value.(*dto.User),&posts[i].UserInfo)
//	}
//	return posts,nil
//}
//func (p postService) Post(ctx *gin.Context,post *dto.Post) (*dto.Post,error) {
//	jwtService := NewJWTService()
//	if ! jwtService.IsPermit(post.UserId,ctx){
//		return nil,errors.New("Can not access")
//	}
//	if p.DB.Where("ID=?",post.CategoryId).Find(&dto.Category{}).RowsAffected ==0{
//		return nil,errors.New(fmt.Sprintf("Not exists category with id %v",post.CategoryId))
//	}
//
//	post.Category = p.DB.Where("ID=?",post.CategoryId).Find(&dto.Category{}).Value.(*dto.Category)
//	mapstructure.Decode(p.DB.Where("ID=?",post.UserId).Find(&dto.User{}).Value.(*dto.User),&post.UserInfo)
//
//
//	if err := p.DB.Create(&post).Error;err!=nil{
//		return nil, err
//	}
//	return post,nil
//}

func NewPostService() PostService {

	if postInstance == nil{
		postOnce.Do(
			func() {
				postInstance = &postService{DB: common.ConnectData()}
			})
	}
	return postInstance
}