package controller

import (
	"com.phh/start-web/model/result"
	"com.phh/start-web/service/sysservice"
	"com.phh/start-web/util"
	"github.com/cristalhq/jwt/v4"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"time"
)

var UserSet = wire.NewSet(wire.Struct(new(UserController), "UserService", "Jwt"))

type UserController struct {
	Ctx         *gin.Context
	UserService *sysservice.UserService
	Jwt         *util.JwtHelper
}

// GetBy : http://localhost:8080/user?id=1
func (a *UserController) GetBy(c *gin.Context) {
	id := c.GetInt("id")
	result.OkData(a.UserService.GetById(id), c)
}

// GetToken 登录
func (a *UserController) GetToken(c *gin.Context) {
	var username = c.Query("username")
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
	result.OkData(token.String(), c)
}
