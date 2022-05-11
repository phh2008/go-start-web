package middleware

import (
	"com.phh/start-web/auth"
	"com.phh/start-web/model"
	"com.phh/start-web/util"
	"fmt"
	"github.com/kataras/iris/v12"
)

func CasbinHandle(casbinHelper *auth.CasbinHelper) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		// 获取当前用户信息
		user, ok := ctx.Values().Get("user").(util.UserClaims)
		if !ok {
			ctx.JSON(model.NewResult("1000", "登录无效", nil))
			return
		}
		// TODO 当前用户的角色
		user.Rule = "admin"
		request := ctx.Request()
		uri := request.URL.Path
		method := request.Method
		fmt.Println("rule:", user.Rule, "uri:", uri, "method:", method)
		enforcer := casbinHelper.GetEnforcer()
		if ok, _ = enforcer.Enforce(user.Rule, uri, method); ok {
			ctx.Next()
		} else {
			ctx.JSON(model.NewResult("1001", "无权限", nil))
		}
	}
}
