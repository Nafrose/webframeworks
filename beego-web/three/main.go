package main

import (
	_ "three/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

