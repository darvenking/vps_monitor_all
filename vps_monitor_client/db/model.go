package db

import (
	"gorm.io/gorm"
)

func init() {
	// 迁移 schema
	err := DB.AutoMigrate(&SiteInfo{})
	if err != nil {
		panic("数据库迁移失败！")
	}
}

var SiteInfoDB = DB.Model(&SiteInfo{})

type SiteInfo struct {
	gorm.Model
	Stock bool
	URL   string
	Name  string
}
