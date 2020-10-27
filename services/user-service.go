package services

import (
	"demo/common"
	"demo/models"
	"demo/models/dto"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"sync"
)
var userInit sync.Once
var userInstance *userService



type UserService interface {
	//PermitAccess(ctx *gin.Context,userId uint) bool
	//FindOne(condition interface{}) *dto.User
	//GetById(ctx *gin.Context) (*models.UserInfo,error)
	//GetAll(ctx *gin.Context) ([]models.UserInfo,error)
	//Register(user *dto.User) (*dto.User,error)
	//Update(ctx *gin.Context) (*dto.User,error)

	ConvertUserToUserInfo(user *dto.User) *models.UserInfo

}

type userService struct {
	 DB *gorm.DB
}

func (u userService) ConvertUserToUserInfo(user *dto.User) *models.UserInfo {
	var userInfo models.UserInfo
	mapstructure.Decode(user,&userInfo)
	return &userInfo
}

//func (u userService) PermitAccess(ctx *gin.Context,userId uint) bool {
//	jwtService := NewJWTService()
//	claims := jwtService.GetClaims(ctx)
//	return claims.ID == userId || claims.IsAdmin
//}
//
//func (u userService) FindOne(condition interface{}) *dto.User {
//	var user dto.User
//	u.DB.Where(condition).Find(&user)
//	return &user
//}
//
//func (u userService) GetById(ctx *gin.Context) (*models.UserInfo, error) {
//
//	id,_ := strconv.Atoi(ctx.Query("id"))
//
//	var user dto.User
//
//	if err := u.DB.Where("ID=?",uint(id)).Find(&user).Error;err!=nil{
//		return nil,err
//	}
//
//	var user_infor models.UserInfo
//	mapstructure.Decode(user,&user_infor)
//	return &user_infor, nil
//
//}
//
//func (u userService) Update(ctx *gin.Context) (*dto.User,error) {
//
//	//Get id from Param
//
//
//	//find user by id
//	//findUser := u.FindById(uint(id))
//	//if findUser==nil{
//	//	return nil,errors.New("Not exist user")
//	//}
//
//	//bind user from client request
//	var user dto.User
//	if err := ctx.ShouldBindJSON(&user);err!=nil{
//		return nil,err
//	}
//	//Check role user update
//	// just user or admin can update
//	jwtService := NewJWTService()
//	if ! jwtService.IsPermit(user.ID,ctx){
//		return nil, errors.New("Can not access")
//	}
//
//
//	//find user by id
//	if err := u.DB.Where("ID=?",user.ID).Find(&dto.User{}).Error;err!=nil{
//		return nil, err
//	}
//	//check email exist
//	if err := u.DB.Where(&dto.User{Email: user.Email}).Find(&dto.User{}).Error;err==nil{
//		return nil,errors.New("Email already in database")
//	}
//
//
//	if err := u.DB.Save(&user).Error;err!=nil{
//		return nil,err
//	}
//	return &user,nil
//}
//
//func (u userService) Register(user *dto.User) (*dto.User,error) {
//
//	//check email exist
//	if err := u.DB.Where(&dto.User{Email: user.Email}).Find(&dto.User{}).Error;err==nil{
//		return nil,errors.New("Email already in database")
//	}
//
//	if err := u.DB.Create(&user).Error;err!=nil{
//		return nil,err
//	}
//	return user,nil
//}
//
//func (u userService) GetAll(ctx *gin.Context) ([]models.UserInfo,error) {
//
//	jwtService := NewJWTService()
//	if ! jwtService.IsPermit(0,ctx){
//		return nil, errors.New("Can not access")
//	}
//
//	var users []dto.User
//	u.DB.Find(&users)
//	var user_infors []models.UserInfo
//	mapstructure.Decode(users,&user_infors)
//	return user_infors,nil
//}


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