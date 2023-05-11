package main

import (
	"fmt"
	"web/core"
	_ "web/docs"
	"web/flag"
	"web/global"
	"web/routers"
)

// @title web API文档
// @version 1.0
// @description web API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	core.InitConf()
	fmt.Println(global.Config)
	global.DB = core.InitGorm()
	// fmt.Println(global.DB)
	global.Log = core.InitLogger()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// 连接redis
	global.Redis = core.ConnectRedis()
	// 连接es
	global.ESclient = core.EsConnect()

	router := routers.InitRouter()
	err := router.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
