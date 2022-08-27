package routers

import (
	"beego-rbac/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/RoleList", &controllers.MainController{}, "get:RoleList")
	beego.Router("/RoleListJson", &controllers.MainController{}, "get:RoleListJson")
	beego.Router("/RoleAdd", &controllers.MainController{}, "post:RoleAdd")
	beego.Router("/RoleDelete", &controllers.MainController{}, "post:RoleDelete")
	beego.Router("/NodeList", &controllers.MainController{}, "get:NodeList")
	beego.Router("/NodeListJson", &controllers.MainController{}, "get:NodeListJson")
	beego.Router("/NodeAdd", &controllers.MainController{}, "post:NodeAdd")
	beego.Router("/NodeDelete", &controllers.MainController{}, "post:NodeDelete")
	beego.Router("/UserList", &controllers.MainController{}, "get:UserList")
	beego.Router("/UserListJson", &controllers.MainController{}, "get:UserListJson")
	beego.Router("/UserAdd", &controllers.MainController{}, "post:UserAdd")
	beego.Router("/UserDelete", &controllers.MainController{}, "post:UserDelete")
	beego.Router("/AccessListJson", &controllers.MainController{}, "get:AccessListJson")
	beego.Router("/AccessAdd", &controllers.MainController{}, "post:AccessAdd")
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/LoginSubmit", &controllers.LoginController{}, "post:LoginSubmit")
	beego.Router("/exit", &controllers.LoginController{}, "get:Quit")
	beego.Router("/404", &controllers.LoginController{}, "get:ErrorPage")
}
