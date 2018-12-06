package model

import (
	"github.com/jinzhu/gorm"
	"GinWebBase/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// 开启数据库连接
func OpenDB(dialect string, path string) {
	var err error

	DB, err = gorm.Open(dialect, path)
	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(true)
}

// 初始化数据库
func InitDB(conf *config.Config) *gorm.DB {
	OpenDB(
		conf.DataBase.Type,
		conf.DataBase.ConnectStr,
	)

	if conf.DataBase.Init {
		DB.CreateTable(&Users{})
	}
	DB.AutoMigrate(&Users{})

	return DB
}
