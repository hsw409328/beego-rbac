package main

import (
	_ "beego-rbac/routers"
	beego "github.com/beego/beego/v2/adapter"
)

func main() {
	beego.SetStaticPath("/layui", "/static/layui")
	beego.BConfig.WebConfig.TemplateLeft = "[[["
	beego.BConfig.WebConfig.TemplateRight = "]]]"
	beego.Run()
}
