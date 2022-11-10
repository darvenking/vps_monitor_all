package controller

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/gins"
	"github.com/gogf/gf/v2/net/ghttp"
	"vps_monitor/internal/model"
	"vps_monitor/utility/res"
)

func Submit(r *ghttp.Request) {
	msg := "商家：" + r.Get("name").String() + "\nUrl：" + r.Get("url").String() + "\n价格：" + r.Get("price").String() + "\n商品名：" + r.Get("productName").String()

	param := gmap.HashMap{}
	param.Set("chat_id", "1810124852")
	param.Set("text", msg)

	url := "https://tg.guoer.cc/bot1974237821:AAH0mWq2JeHz_GkK0Mb8xLIZyGCyekdOUxE/sendMessage"

	result, err := gins.HttpClient().Post(r.GetCtx(), url, param)
	if err != nil {
		return
	}
	println(result)
	res.Success(r, "ok")
}

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
