package service

import "github.com/gogf/gf/v2/net/ghttp"

type MiddlewareService struct{}

var middleware = MiddlewareService{}

func Middleware() *MiddlewareService {
	return &middleware
}

func (s *MiddlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *MiddlewareService) Auth(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
