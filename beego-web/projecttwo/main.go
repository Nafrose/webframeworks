package main

import beego "github.com/beego/beego/v2/server/web"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("Hello world from project two with beego!")
}

func main(){
	beego.Router("/", &MainController{})
	beego.Run()
}
