package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"vps_monitor/internal/service"
	"vps_monitor/utility/res"
)

func Login(r *ghttp.Request) {
	tokenString, _ := service.Auth().LoginHandler(r.GetCtx())
	res.Success(r, tokenString)
}

func Logout(r *ghttp.Request) {
	service.Auth().LogoutHandler(r.GetCtx())
	res.Success(r, "")
}

func Test(r *ghttp.Request) {
	res.Success(r, "test"+r.Get("userId").String())
}
