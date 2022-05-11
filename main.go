package main

import (
	"com.phh/start-web/config"
	"com.phh/start-web/web/router"
	"flag"
	"github.com/kataras/iris/v12"
)

func main() {
	var configFolder string
	flag.StringVar(&configFolder, "config", "./config", "指定配置文件目录，e.g. -config ./config")
	flag.Parse()
	if configFolder == "" {
		configFolder = "./config"
	}
	app := iris.New()
	app.Logger().SetLevel("debug")
	appCtx := BuildApp(config.ConfigFolder(configFolder))
	// 初始化 casbin
	_ = appCtx.Casbin.GetEnforcer()
	router.Register(app, appCtx)
	_ = app.Run(
		// 启动端口
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
