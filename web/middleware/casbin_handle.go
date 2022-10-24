package middleware

import (
	"com.phh/start-web/auth"
	"com.phh/start-web/model/result"
	"com.phh/start-web/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinHandle(casbinHelper *auth.CasbinHelper) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 获取当前用户信息
		u, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusOK, result.NewResult("1000", "登录无效", nil))
			return
		}
		user, ok := u.(util.UserClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusOK, result.NewResult("1000", "登录无效", nil))
			return
		}
		// TODO 当前用户的角色
		user.Rule = "admin"
		request := ctx.Request
		uri := request.URL.Path
		method := request.Method
		fmt.Println("rule:", user.Rule, "uri:", uri, "method:", method)
		enforcer := casbinHelper.GetEnforcer()
		if ok, _ = enforcer.Enforce(user.Rule, uri, method); ok {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusOK, result.NewResult("1001", "无权限", nil))
		}
	}
}
