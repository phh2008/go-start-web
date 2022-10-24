package controller

import (
	"com.phh/start-web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var AuthSet = wire.NewSet(wire.Struct(new(AuthController), "*"))

type AuthController struct {
	Casbin *auth.CasbinHelper
}

// Hello 测试 api
func (a *AuthController) Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "请求成功：success")
}

// AddAuth 测试添加权限
func (a *AuthController) AddAuth(ctx *gin.Context) {
	if ok, _ := a.Casbin.GetEnforcer().AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
		ctx.String(http.StatusOK, "Policy已经存在")
	} else {
		ctx.String(http.StatusOK, "增加成功")
	}
}

// DelAuth 测试删除权限
func (a *AuthController) DelAuth(ctx *gin.Context) {
	if ok, _ := a.Casbin.GetEnforcer().RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
		ctx.String(http.StatusOK, "Policy不存在")
	} else {
		ctx.String(http.StatusOK, "删除成功")
	}
}
