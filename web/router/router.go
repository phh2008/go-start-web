package router

import (
	"com.phh/start-web/app"
	"com.phh/start-web/web/middleware"
	"github.com/gin-gonic/gin"
)

// Register 注册路由
func Register(app *gin.Engine, appCtx *app.AppContext) {
	// 中间件
	jwtAuth := middleware.JwtMiddleware(appCtx.Jwt)
	casbin := middleware.CasbinHandle(appCtx.Casbin)

	// 用户
	userPath := app.Group("/user")
	{
		userPath.GET("/user", appCtx.UserApi.GetBy)
		userPath.GET("/token", appCtx.UserApi.GetToken)
	}

	// 订单
	orderPath := app.Group("/order")
	orderPath.Use(jwtAuth)
	orderPath.Use(casbin)
	{
		orderPath.GET("/{id:int}", appCtx.OrderApi.GetById)
		orderPath.GET("/query", appCtx.OrderApi.Query)
	}

	// auth
	v1 := app.Group("/api/v1")
	v1.Use(jwtAuth)
	v1.Use(casbin)
	{
		v1.GET("/hello", appCtx.AuthApi.Hello)
		v1.GET("/addAuth", appCtx.AuthApi.AddAuth)
		v1.GET("/delAuth", appCtx.AuthApi.DelAuth)
	}

}
