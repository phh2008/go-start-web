package controller

import (
	"com.phh/start-web/entity"
	"com.phh/start-web/service"
	"com.phh/start-web/util"
	"github.com/cristalhq/jwt/v4"
	"github.com/google/wire"
	"github.com/kataras/iris/v12"
	"time"
)

var UserSet = wire.NewSet(wire.Struct(new(UserController), "UserService", "Jwt"))

type UserController struct {
	Ctx         iris.Context
	UserService *service.UserService
	Jwt         *util.JwtHelper
}

// GetBy : http://localhost:8080/user?id=1
func (a *UserController) GetBy(id int) entity.User {
	return a.UserService.GetById(id)
}

// GetToken 登录
func (a *UserController) GetToken() string {
	var username = a.Ctx.URLParam("username")
	var userClaims = util.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:      "1000",
			Subject: username,
			// 有效期 30天
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24 * 30)},
		},
		Phone: "18975391618",
	}
	token, _ := a.Jwt.CreateToken(userClaims)
	return token.String()
}
