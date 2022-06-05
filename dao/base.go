package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-demo/model"
)

var db *gorm.DB

func init() {
	Db, err := gorm.Open(mysql.Open("dy:douyin@tcp(127.0.0.1:3306)/simple_demo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = Db
	db.AutoMigrate(&model.User{}, &model.UserFollow{}, &model.Video{}, &model.UserFavorite{}, &model.Comment{})
}
