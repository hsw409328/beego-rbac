package controllers

import (
	"beego-rbac/models"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

var (
	leftTreeResultMap = make(map[int][]orm.Params)
)

type BaseController struct {
	beego.Controller
}

// 设置当前登录用户的可以查看的页面权限
func (this *BaseController) SessionRbacNav(id int) {
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

	leftTreeResultMap[id] = []orm.Params{}
	tmpResult := this.TreeNodeRecursion(id, accessResult, 0)
	this.SetSession("LeftNavResult", tmpResult)
	delete(leftTreeResultMap, id)
}

// 设置当前登录用户的所有权限
func (this *BaseController) SessionRbacAll(id int) {
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
		" where t1.status=1 and t.role_id='" + strconv.Itoa(roleResult.Id) + "' ").Values(&accessResult)

	leftTreeResultMap[id] = []orm.Params{}
	tmpResult := this.TreeNodeRecursion(id, accessResult, 0)
	this.SetSession("LoginUserAllRBACResult", tmpResult)
	delete(leftTreeResultMap, id)
}

func (this *BaseController) TreeNodeRecursion(userId int, data []orm.Params, pid int) []orm.Params {
	for _, v := range data {
		tmpPid, _ := strconv.Atoi(v["pid"].(string))
		if tmpPid == pid {
			leftTreeResultMap[userId] = append(leftTreeResultMap[userId], v)
			tmpId, _ := strconv.Atoi(v["id"].(string))
			this.TreeNodeRecursion(userId, data, tmpId)
		}
	}
	return leftTreeResultMap[userId]
}

func (this *BaseController) CheckLogin() (interface{}, bool) {
	userData := this.GetSession("LoginUser")
	if userData == nil {
		return nil, false
	}
	return userData, true
}
