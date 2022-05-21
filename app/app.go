package app

import (
	"com.phh/start-web/auth"
	"com.phh/start-web/pkg/config"
	"com.phh/start-web/pkg/logger"
	"com.phh/start-web/util"
	"com.phh/start-web/web/controller"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var AppSet = wire.NewSet(wire.Struct(new(AppContext), "*"))

type AppContext struct {
	Config   *config.Config
	DB       *gorm.DB
	Jwt      *util.JwtHelper
	Casbin   *auth.CasbinHelper
	UserApi  *controller.UserController
	OrderApi *controller.OrderController
	AuthApi  *controller.AuthController
	Logger   *logger.Logger
}
