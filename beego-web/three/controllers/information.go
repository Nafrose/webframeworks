package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type InfoController struct {
	beego.Controller
}

func (c *InfoController) Get() {

	c.TplName = "information.tpl"
}
