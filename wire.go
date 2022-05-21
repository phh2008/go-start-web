//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"com.phh/start-web/app"
	"com.phh/start-web/auth"
	"com.phh/start-web/dao"
	"com.phh/start-web/pkg/config"
	"com.phh/start-web/pkg/logger"
	"com.phh/start-web/service"
	"com.phh/start-web/util"
	"com.phh/start-web/web/controller"
	"github.com/google/wire"
)

func BuildApp(configFolder config.ConfigFolder) *app.AppContext {
	wire.Build(
		config.ConfigSet,
		util.InitDB,
		util.NewJwtHelper,
		auth.CasbinSet,
		dao.DaoSet,
		service.ServiceSet,
		controller.ControllerSet,
		logger.LoggerSet,
		app.AppSet,
	)
	return new(app.AppContext)
}
