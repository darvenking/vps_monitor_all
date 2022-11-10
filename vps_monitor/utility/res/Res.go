package res

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

func Success(r *ghttp.Request, data interface{}) {
	r.Response.WriteJsonExit(&ghttp.DefaultHandlerResponse{
		Data: data,
		Code: http.StatusOK,
	})
}

func Fail(r *ghttp.Request, msg string) {
	r.Response.WriteJsonExit(&ghttp.DefaultHandlerResponse{
		Message: msg,
		Code:    http.StatusForbidden,
	})
}
