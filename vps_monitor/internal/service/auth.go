package service

import (
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
	"vps_monitor/internal/model"
	"vps_monitor/internal/param"
	"vps_monitor/utility/res"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "auth zone",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 30,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "userId",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,   //根据登录信息对用户进行身份验证的回调函数
		Unauthorized:    Unauthorized,    //处理不进行授权的逻辑
		PayloadFunc:     PayloadFunc,     //登录期间的设置私有载荷的函数，默认设置Authenticator函数回调的所有内容
		IdentityHandler: IdentityHandler, //解析并设置用户身份信息，并设置身份信息至每次请求中
		CacheAdapter:    gcache.NewAdapterRedis(g.Redis()),
	})
	authService = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	r := g.RequestFromCtx(ctx)
	var reqParam = &param.UserLoginParam{}
	if err := r.Parse(&reqParam); err != nil {
		res.Fail(r, err.Error())
	}
	var userInfo *model.UserInfo
	model.GetUserInfoDB().Where("user_name = ?", reqParam.UserName).First(&userInfo)
	if userInfo == nil {
		res.Fail(r, "用户名不存在")
	}
	if userInfo.PassWord != reqParam.PassWord {
		res.Fail(r, "密码不正确")
	}
	return g.Map{
		"userId":   userInfo.ID,
		"username": userInfo.UserName,
	}, nil
}
