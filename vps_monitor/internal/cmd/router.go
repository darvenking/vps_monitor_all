package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"vps_monitor/internal/controller"
	"vps_monitor/internal/service"
)

func route(ctx context.Context, parser *gcmd.Parser) error {
	s := g.Server()
	s.BindMiddlewareDefault(ghttp.MiddlewareHandlerResponse)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.POST("/login", controller.Login)
		group.GET("/seller", controller.SellerList)
		group.POST("/plist", controller.Plist)
		group.POST("/submit", controller.Submit)
	})

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.POST("/logout", controller.Logout)
		group.POST("/audit", controller.Audit)

	})
	s.Run()
	return nil
}
