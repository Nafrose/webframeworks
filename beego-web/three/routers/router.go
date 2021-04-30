package routers

import (
	"github.com/beego/beego/v2/server/web"
	"three/controllers"
)

func init() {
    web.Router("/", &controllers.MainController{})
	web.Router("/info", &controllers.InfoController{})
	web.SetStaticPath("/static/dl1/test.txt", "/static/dl1/test.txt")
	//web.SetStaticPath("/dl1/test.txt", "dl1/test.txt")
}
