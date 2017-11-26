package controllers

import (
	"beego-project/models"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	_, err := this.CheckLogin()
	if err {
		this.Redirect("/", 302)
		this.StopRun()
	}
	this.TplName = "login.html"
}

func (this *LoginController) LoginSubmit() {
	username := this.GetString("username")
	password := this.GetString("password")
	result := models.RbacUser{}
	err := M.Object().QueryTable(userModel).Filter("username", username).Filter("password", password).One(&result)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "991", "error_msg": "登录失败，用户名或密码错误"}
		this.ServeJSON()
	} else {
		this.SetSession("LoginUser", result)
		this.SessionRbac(result.Id)
		this.Data["json"] = map[string]string{"code": "0", "error_msg": "登录成功，跳转中..."}
		this.ServeJSON()
	}
	this.StopRun()
}

func (this *LoginController) Quit() {
	this.DestroySession()
	this.Redirect("/login", 302)
	this.StopRun()
}

func (this *LoginController) ErrorPage() {
	this.TplName = "error.html"
}
