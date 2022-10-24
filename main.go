package main

import (
	"com.phh/start-web/pkg/config"
	"com.phh/start-web/pkg/global"
	"com.phh/start-web/web/middleware"
	"com.phh/start-web/web/router"
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	// 命令行参数
	var configFolder string
	flag.StringVar(&configFolder, "config", "./config", "指定配置文件目录，e.g. -config ./config")
	flag.Parse()
	if configFolder == "" {
		configFolder = "./config"
	}
	// wire
	appCtx := BuildApp(config.ConfigFolder(configFolder))
	global.Log = appCtx.Logger
	if err := appCtx.Config.Viper.Unmarshal(&global.Profile); err != nil {
		panic(err)
	}
	global.Log.Infof("----------------- start -----------------")
	// 初始化 casbin
	_ = appCtx.Casbin.GetEnforcer()
	// gin
	app := gin.Default()
	app.Use(middleware.GinRecovery(true))
	app.Use(middleware.Translations())
	router.Register(app, appCtx)
	_ = app.Run(":8088")
}
