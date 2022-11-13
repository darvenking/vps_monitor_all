package db

import (
	"gorm.io/gorm"
)

func init() {
	// 迁移 schema
	err := DB.AutoMigrate(&SiteInfo{}, &SiteConfig{}, &SellerInfo{})
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
	ConfigId    uint
}

func GetSubmitSiteDB() *gorm.DB {
	return DB.Model(&SiteConfig{})
}

type SubmitSite struct {
	gorm.Model
	URL string
	//处理状态，1未处理 2已处理
	Status int `gorm:"default:1"`
}

func GetSiteConfigDB() *gorm.DB {
	return DB.Model(&SiteConfig{})
}

type SiteConfig struct {
	gorm.Model
	URL         string
	NoStockFlag string
	NameFlag    string
	PriceFlag   string
	Cookies     string
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
