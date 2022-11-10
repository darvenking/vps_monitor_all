package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"vps_monitor/internal/model"
	"vps_monitor/utility/res"
)

func Add(r *ghttp.Request) {
	info := &model.SiteInfo{
		URL:  r.Get("url").String(),
		Name: r.Get("name").String(),
	}
	model.GetSiteInfoDB().Save(info)
	res.Success(r, "ok")
}

func Plist(r *ghttp.Request) {
	id := r.Get("id").Int()
	stock := r.Get("stock").Int()
	var list []model.SiteInfo
	siteInfoDB := model.GetSiteInfoDB()
	if id != 0 {
		siteInfoDB.Where("seller_id = ?", id)
	}
	if stock != 0 {
		siteInfoDB.Where("stock = ?", stock == 2)
	}
	siteInfoDB.Find(&list)
	res.Success(r, list)
}
