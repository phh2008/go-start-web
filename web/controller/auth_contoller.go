package controller

import (
	"com.phh/start-web/auth"
	"github.com/google/wire"
	"github.com/kataras/iris/v12"
)

var AuthSet = wire.NewSet(wire.Struct(new(AuthController), "*"))

type AuthController struct {
	Casbin *auth.CasbinHelper
}

// Hello 测试 api
func (a *AuthController) Hello(ctx iris.Context) {
	ctx.WriteString("请求成功：success")
}

// AddAuth 测试添加权限
func (a *AuthController) AddAuth(ctx iris.Context) {
	if ok, _ := a.Casbin.GetEnforcer().AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
		ctx.WriteString("Policy已经存在")
	} else {
		ctx.WriteString("增加成功")
	}
}

// DelAuth 测试删除权限
func (a *AuthController) DelAuth(ctx iris.Context) {
	if ok, _ := a.Casbin.GetEnforcer().RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
		ctx.WriteString("Policy不存在")
	} else {
		ctx.WriteString("删除成功")
	}
}
