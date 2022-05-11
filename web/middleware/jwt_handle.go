package middleware

import (
	"com.phh/start-web/model"
	"com.phh/start-web/util"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"time"
)

func JwtMiddleware(jwtHelper *util.JwtHelper) func(ctx iris.Context) {
	jwtMiddleware := func(ctx iris.Context) {
		// before
		token := ctx.GetHeader("token")
		if token == "" {
			token = ctx.URLParam("token")
		}
		if token == "" {
			ctx.JSON(model.NewResult("1000", "未登录", nil))
			return
		}
		jwtToken, err := jwtHelper.VerifyToken(token)
		if err != nil {
			ctx.JSON(model.NewResult("1000", "登录无效", err))
			//ctx.StopWithJSON(200,model.NewResult("1000", "登录无效", err))
			return
		}
		var user util.UserClaims
		json.Unmarshal(jwtToken.Claims(), &user)
		if !user.IsValidExpiresAt(time.Now()) {
			ctx.JSON(model.NewResult("1000", "登录过期", err))
			return
		}
		ctx.Values().Set("user", user)
		ctx.Next()
		// after
	}
	return jwtMiddleware
}
