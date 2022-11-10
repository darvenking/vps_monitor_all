package model

/**
文档: https://gorm.io/zh_CN/docs/query.html
*/

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"vps_monitor/utility/cfg"
)

var (
	DB = GetDB()
)

func init() {
	// 迁移 schema
	err := DB.AutoMigrate(&SiteInfo{}, &UserInfo{}, &SellerInfo{})
	if err != nil {
		panic("数据库迁移失败！")
	}
}

func GetDB() *gorm.DB {
	host := cfg.Get("database.host")
	user := cfg.Get("database.user")
	password := cfg.Get("database.password")
	dbname := cfg.Get("database.dbname")
	port := cfg.GetStr("database.port")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	glog.Info(gctx.New(), "数据连接初始化完成")
	return db
}
