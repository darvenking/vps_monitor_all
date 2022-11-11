package db

import (
	"gorm.io/gorm"
)

func init() {
	// 迁移 schema
	err := DB.AutoMigrate(&SiteInfo{}, &SitePre{})
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

func GetSitePreDB() *gorm.DB {
	return DB.Model(&SitePre{})
}

type SitePre struct {
	gorm.Model
	URL         string
	NoStockFlag string
	NameFlag    string
	PriceFlag   string
	//处理状态，1待审核，2已审核待爬虫处理，3爬虫已处理完毕
	Status int `gorm:"default:1"`
}

func GetSellerInfoDB() *gorm.DB {
	return DB.Model(&SellerInfo{})
}

type SellerInfo struct {
	gorm.Model
	SellerName  string
	Description string
	Status      int `gorm:"default:1"`
}
