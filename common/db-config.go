package common

import (
	"demo/models/dto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
)
var once sync.Once
var DB *gorm.DB =nil
func ConnectData() *gorm.DB {
	if DB==nil{
		once.Do(
			func() {
				db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=demo password=123 sslmode=disable")

				if err != nil {
					panic("Fail to connect database ")
				}

				//migrate data
				db.AutoMigrate(&dto.User{})
				db.AutoMigrate(&dto.Category{})
				db.AutoMigrate(&dto.Post{})
				DB = db
			},
		)
	}
	return DB
}

