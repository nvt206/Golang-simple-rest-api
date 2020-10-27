package services

import (
	"demo/common"
	"github.com/jinzhu/gorm"
	"sync"
)

var baseInstance *baseService
var baseOnce sync.Once

type BaseService interface {
	GetAll(receive interface{}) error
	GetOne(receive interface{}, where string,args ...interface{}) error
	GetList(receive interface{}, where string, args ...interface{}) error
	Create(req interface{}) error
}

type baseService struct {
	DB *gorm.DB
}

func (b baseService) GetOne(receive interface{}, where string,args ...interface{}) error {
	return b.DB.Debug().Where(where, args...).Find(receive).Error
}


func (b baseService) GetList(receive interface{}, where string, args ...interface{}) error {
	return b.DB.Debug().Where(where,args...).Find(receive).Error
}

func (b baseService) Create(req interface{}) error {
	return b.DB.Create(req).Error
}

func (b baseService) GetAll(receive interface{}) error {
	return b.DB.Find(receive).Error

}

func NewBaseService() BaseService{
	if baseInstance == nil{
		baseOnce.Do(func() {
			baseInstance = &baseService{DB: common.ConnectData()}
		})
	}
	return baseInstance
}
