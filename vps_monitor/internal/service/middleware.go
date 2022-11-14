package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"vps_monitor/utility/cfg"
	"vps_monitor/utility/res"
)

type middlewareService struct{}

func MiddlewareInstance() *middlewareService {
	return &middlewareService{}
}

func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *middlewareService) Auth(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

// SimpleAuthenticator 简单的自定义授权校验
func (s *middlewareService) SimpleAuthenticator(r *ghttp.Request) {
	secret := r.GetHeader("secret")
	if secret == "" || cfg.GetStr("jwt.secret") != secret {
		res.Fail(r, "授权码不存在")
	}
	r.Middleware.Next()
}
