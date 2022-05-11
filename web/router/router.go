package router

import (
	"com.phh/start-web/app"
	"com.phh/start-web/web/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// Register 注册路由
func Register(app *iris.Application, appCtx *app.AppContext) {
	// 中间件
	jwtAuth := middleware.JwtMiddleware(appCtx.Jwt)
	casbin := middleware.CasbinHandle(appCtx.Casbin)

	// 用户
	userPath := app.Party("/user")
	user := mvc.New(userPath)
	user.Handle(appCtx.UserApi)

	// 订单
	orderPath := app.Party("/order")
	orderPath.Use(jwtAuth)
	orderPath.Use(casbin)
	orderPath.Handle(iris.MethodGet, "/{id:int}", appCtx.OrderApi.GetById)
	orderPath.Get("/query", appCtx.OrderApi.Query)

	// auth
	v1 := app.Party("/api/v1")
	v1.Use(jwtAuth)
	v1.Use(casbin)
	v1.Get("/hello", appCtx.AuthApi.Hello)
	v1.Get("/addAuth", appCtx.AuthApi.AddAuth)
	v1.Get("/delAuth", appCtx.AuthApi.DelAuth)

}
