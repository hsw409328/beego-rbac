package controllers

import (
	"beego-rbac/models"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type BaseController struct {
	beego.Controller
}

var (
	leftTreeResult = []orm.Params{}
)

func (this *BaseController) SessionRbac(id int) {
	//根据用户ID查询出角色
	result := models.RbacRoleUser{UserId: id}
	M.Object().Read(&result)
	//查询出角色信息
	roleResult := models.RbacRole{Id: result.RoleId}
	M.Object().Read(&roleResult)
	this.SetSession("RoleInfo", roleResult)
	//根据角色查询出权限
	var accessResult []orm.Params
	M.Object().Raw("select t1.id,t1.name,t1.title,t1.sort,t1.pid,t1.level,t.role_id" +
		" from rbac_access t inner join rbac_node t1" +
		" on t.node_id=t1.id " +
		" where t1.status=1 and t.role_id='" + strconv.Itoa(roleResult.Id) + "' and t1.is_show=1 ").Values(&accessResult)

	leftTreeResult = []orm.Params{}
	tmpResult := this.TreeNodeRecursion(accessResult, 0)
	this.SetSession("LeftNavResult", tmpResult)
}

func (this *BaseController) TreeNodeRecursion(data []orm.Params, pid int) []orm.Params {
	for _, v := range data {
		tmpPid, _ := strconv.Atoi(v["pid"].(string))
		if tmpPid == pid {
			leftTreeResult = append(leftTreeResult, v)
			tmpId, _ := strconv.Atoi(v["id"].(string))
			this.TreeNodeRecursion(data, tmpId)
		}
	}
	return leftTreeResult
}

func (this *BaseController) CheckRbac(module string) bool {
	tmpResult := this.GetSession("LeftNavResult")
	if tmpResult == nil {
		return false
	}
	returnResult := false
	for _, v := range tmpResult.([]orm.Params) {
		if v["name"].(string) == module {
			returnResult = true
			break
		}
	}
	return returnResult
}

func (this *BaseController) CheckLogin() (interface{}, bool) {
	userData := this.GetSession("LoginUser")
	if userData == nil {
		return nil, false
	}
	return userData, true
}

func (this *BaseController) CheckRbacAll(module string) {
	info, err := this.CheckLogin()
	infoResult := info.(models.RbacUser)
	if !err {
		this.Redirect("/login", 302)
		this.StopRun()
	}
	u := beego.AppConfig.String("superadmin")
	if infoResult.Username == u {
		return
	}
	err = this.CheckRbac(module)
	if !err {
		this.Redirect("/404", 302)
		this.StopRun()
	}
}
