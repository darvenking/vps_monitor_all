package db

/**
文档: https://gorm.io/zh_CN/docs/query.html
*/

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"vps_monitor_client/util"
)

var (
	DB = GetDB()
)

func GetDB() *gorm.DB {

	host := util.GetCfgStr("database.host")
	user := util.GetCfgStr("database.user")
	password := util.GetCfgStr("database.password")
	dbname := util.GetCfgStr("database.dbname")
	port := util.GetCfgStr("database.port")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	println("数据连接初始化完成")
	return db
}
