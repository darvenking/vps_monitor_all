package db

/**
文档: https://gorm.io/zh_CN/docs/query.html
*/

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB = GetDB()
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
