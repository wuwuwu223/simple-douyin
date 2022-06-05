package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"simple-demo/global"
	"simple-demo/model"
	"time"
)

var db *gorm.DB

func InitDb() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true", global.Config.Mysql.User, global.Config.Mysql.Password, global.Config.Mysql.Host, global.Config.Mysql.Port, global.Config.Mysql.Database)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		PrepareStmt:            true, //开启缓存
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db = Db
	db.AutoMigrate(&model.User{}, &model.UserFollow{}, &model.Video{}, &model.UserFavorite{}, &model.Comment{})
}
