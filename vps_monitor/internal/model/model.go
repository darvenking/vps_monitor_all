package model

import (
	"gorm.io/gorm"
)

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
}

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
