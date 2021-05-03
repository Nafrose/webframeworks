package main

import (
"github.com/beego/beego/v2/core/logs"
"github.com/beego/beego/v2/server/web"
)

var (
	BuildMode       = "dev"
)

func main() {
	web.AppConfig.Set("BuildMode", BuildMode)

	val, _ := web.AppConfig.String("BuildMode")

	logs.Info("load config BuildMode is", val)

	web.Run()
}


