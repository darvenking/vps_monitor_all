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

func GetSiteInfoDB() *gorm.DB {
	return DB.Model(&SiteInfo{})
}

type SiteInfo struct {
	gorm.Model
	Stock       bool
	NoStockFlag string
	URL         string
	Name        string
	Price       string
	Status      int `gorm:"default:1"`
	SellerId    uint
}
