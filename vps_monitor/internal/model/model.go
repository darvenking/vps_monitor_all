package model

import (
	"gorm.io/gorm"
)

func GetUserInfoDB() *gorm.DB {
	return DB.Model(&UserInfo{})
}

type UserInfo struct {
	gorm.Model
	UserName string
	PassWord string
	Status   int `gorm:"default:1"`
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
	//处理状态，1待爬虫处理，3爬虫已处理完毕
	Status int `gorm:"default:1"`
}

func GetSiteInfoDB() *gorm.DB {
	return DB.Model(&SiteInfo{})
}

type SiteInfo struct {
	gorm.Model
	Stock    bool
	URL      string
	Name     string
	Price    string
	Status   int `gorm:"default:1"`
	SellerId uint
	ConfigId uint
}
