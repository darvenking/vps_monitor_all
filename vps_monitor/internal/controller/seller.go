package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"vps_monitor/internal/model"
	"vps_monitor/utility/res"
)

func SellerList(r *ghttp.Request) {
	var list []model.SellerInfo
	model.GetSellerInfoDB().Where("status = ?", 1).Find(&list) // 根据整型主键查找
	res.Success(r, list)
}
