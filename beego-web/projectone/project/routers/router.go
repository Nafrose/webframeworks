package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"project/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
}
