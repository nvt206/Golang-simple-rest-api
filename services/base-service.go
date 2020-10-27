package services

import (
	"demo/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"sync"
)

var baseInstance *baseService
var baseOnce sync.Once

type BaseService interface {
	GetOne(receive interface{}, where string,args ...interface{}) *gorm.DB
	GetList(receive interface{}, where string, args ...interface{}) *gorm.DB
	Create(req interface{}) *gorm.DB
	Update(req interface{}) *gorm.DB
	Delete(model interface{},where string,args ...interface{}) *gorm.DB
	IsRealUser(ctx *gin.Context,userId uint) bool
	IsAdmin(ctx *gin.Context) bool
}

type baseService struct {
	DB *gorm.DB
}

func (b baseService) IsRealUser(ctx *gin.Context, userId uint) bool {
	jwtService := NewJWTService()
	claims := jwtService.GetClaims(ctx)
	return claims.ID == userId
}
func (b baseService) IsAdmin(ctx *gin.Context) bool {
	jwtService := NewJWTService()
	claims := jwtService.GetClaims(ctx)
	return claims.IsAdmin
}

func (b baseService) Update(req interface{}) *gorm.DB {
	return b.DB.Save(req)
}

func (b baseService) Delete(model interface{}, where string, args ...interface{}) *gorm.DB {
	return b.DB.Where(where,args...).Delete(model)
}
func (b baseService) GetOne(receive interface{}, where string,args ...interface{}) *gorm.DB {
	return b.DB.Debug().Where(where, args...).First(receive)
}
func (b baseService) GetList(receive interface{}, where string, args ...interface{}) *gorm.DB {
	return b.DB.Debug().Where(where,args...).Find(receive)
}
func (b baseService) Create(req interface{}) *gorm.DB {
	return b.DB.Create(req)
}
func NewBaseService() BaseService{
	if baseInstance == nil{
		baseOnce.Do(func() {
			baseInstance = &baseService{DB: common.ConnectData()}
		})
	}
	return baseInstance
}
