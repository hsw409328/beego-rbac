package controllers

import (
	"beego-rbac/models"
	"strconv"
	"github.com/astaxie/beego/orm"
	"strings"
)

var (
	M              = new(models.M)
	roleModel      = new(models.RbacRole)
	nodeModel      = new(models.RbacNode)
	roleuserModel  = new(models.RbacRoleUser)
	userModel      = new(models.RbacUser)
	accessModel    = new(models.RbacAccess)
	resultTreeNode = []models.RbacNode{}
)

type MainController struct {
	BaseController
}

func (this *MainController) JsonEncode(code int, msg string, data interface{}, count int) {
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg, "data": data, "count": count}
	this.ServeJSON()
	this.StopRun()
}

func (this *MainController) Get() {
	info, err := this.CheckLogin()
	if !err {
		this.Redirect("/login", 302)
		this.StopRun()
	}
	this.Data["LeftNavResult"] = this.GetSession("LeftNavResult")
	this.Data["LoginUser"] = info
	this.TplName = "index.html"
}

func (this *MainController) RoleList() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	var result []models.RbacNode
	M.Object().QueryTable(nodeModel).OrderBy("sort").All(&result)
	//每次调用之前，请先清空resultTreeNode,防止重复添加
	resultTreeNode = []models.RbacNode{}
	result = this.TreeNodeRecursion(result, 0)
	this.Data["NodeResult"] = result
	this.TplName = "rbac/role.html"
}

func (this *MainController) RoleListJson() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	var result []models.RbacRole
	n, _ := M.Object().QueryTable("rbac_role").All(&result)
	this.JsonEncode(0, "success", result, int(n))
}

func (this *MainController) RoleAdd() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	id, _ := this.GetInt("id")
	name := this.GetString("name")
	pid, _ := this.GetInt("pid")
	status, _ := this.GetInt("status")
	remark := this.GetString("remark")
	if id != 0 {
		result := models.RbacRole{Id: id}
		if M.Object().Read(&result) == nil {
			result.Remark = remark
			result.Status = status
			result.Pid = pid
			result.Name = name
			if _, err := M.Object().Update(&result); err == nil {
				this.JsonEncode(0, "update success", nil, 0)
			}
			this.JsonEncode(102, "update failed", nil, 0)
		} else {
			this.JsonEncode(102, "update failed", nil, 0)
		}
	}

	result := models.RbacRole{Name: name, Pid: pid, Status: status, Remark: remark}
	_, err := M.Object().Insert(&result)
	if err != nil {
		this.JsonEncode(101, "insert failed", nil, 0)
	}
	this.JsonEncode(0, "success", nil, 0)
}

func (this *MainController) RoleDelete() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	id, _ := this.GetInt("id")
	result := models.RbacRole{Id: id}
	_, err := M.Object().Delete(&result)
	if err != nil {
		this.JsonEncode(103, "delete failed", nil, 0)
	}
	this.JsonEncode(0, "delete success", nil, 0)

}

func (this *MainController) NodeList() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	var result []models.RbacNode
	_, err := M.Object().QueryTable("rbac_node").Filter("level__in", []int{1, 2}).All(&result, "Id", "Name", "Title", "Level")
	if err != nil {

	}
	this.Data["ParentNode"] = result
	this.TplName = "rbac/node.html"
}

func (this *MainController) NodeListJson() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	var result []models.RbacNode
	n, _ := M.Object().QueryTable("rbac_node").OrderBy("sort").All(&result)
	//每次调用之前，请先清空resultTreeNode,防止重复添加
	resultTreeNode = []models.RbacNode{}
	result = this.TreeNodeRecursion(result, 0)
	this.JsonEncode(0, "success", result, int(n))
}

func (this *MainController) TreeNodeRecursion(data []models.RbacNode, pid int) []models.RbacNode {
	for _, v := range data {
		if (v.Pid == pid) {
			resultTreeNode = append(resultTreeNode, v)
			this.TreeNodeRecursion(data, v.Id)
		}
	}
	return resultTreeNode
}

func (this *MainController) NodeAdd() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	this.CheckRbacAll("RoleList")
	id, _ := this.GetInt("id")
	name := this.GetString("name")
	title := this.GetString("title")
	pid, _ := this.GetInt("pid")
	status, _ := this.GetInt("status")
	is_show, _ := this.GetInt("is_show")
	remark := this.GetString("remark")
	sort, _ := this.GetInt("sort")
	level, _ := this.GetInt("level")
	if id != 0 {
		result := models.RbacNode{Id: id}
		if M.Object().Read(&result) == nil {
			result.Remark = remark
			result.Status = status
			result.IsShow = is_show
			result.Pid = pid
			result.Name = name
			result.Level = level
			result.Sort = sort
			result.Title = title
			if _, err := M.Object().Update(&result); err == nil {
				this.JsonEncode(0, "update success", nil, 0)
			}
			this.JsonEncode(102, "update failed", nil, 0)
		} else {
			this.JsonEncode(102, "update failed", nil, 0)
		}
	}

	result := models.RbacNode{Name: name, Pid: pid, Status: status, Remark: remark, Title: title, Sort: sort, Level: level, IsShow: is_show}
	_, err := M.Object().Insert(&result)
	if err != nil {
		this.JsonEncode(101, "insert failed", nil, 0)
	}
	this.JsonEncode(0, "success", nil, 0)
}

func (this *MainController) NodeDelete() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	id, _ := this.GetInt("id")
	result := models.RbacNode{Id: id}
	_, err := M.Object().Delete(&result)
	if err != nil {
		this.JsonEncode(103, "delete failed", nil, 0)
	}
	this.JsonEncode(0, "delete success", nil, 0)

}

func (this *MainController) UserList() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	var result []models.RbacRole
	M.Object().QueryTable(roleModel).Filter("status", 1).All(&result)
	this.Data["RoleResult"] = result
	this.TplName = "rbac/user.html"
}

func (this *MainController) UserListJson() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	var result []orm.Params
	lim, _ := this.GetInt("limit")
	page, _ := this.GetInt("page")
	page = (page - 1) * lim
	//M.Object().QueryTable(userModel).Offset(page).Limit(lim).All(&result)
	M.Object().Raw("select t.id as Id,loginip as Loginip,logintime as Logintime,password as Password," +
		"role_id as RoleId,t.status as Status,user_id as UserId,username as Username,t2.name as Name " +
		" from rbac_user t inner join rbac_role_user t1 " +
		"on t.id=t1.user_id inner join rbac_role t2 on t1.role_id=t2.id " +
		" limit " + strconv.Itoa(page) + "," + strconv.Itoa(lim)).Values(&result)
	count, _ := M.Object().QueryTable(userModel).Count()
	this.JsonEncode(0, "success", result, int(count))
}

func (this *MainController) UserAdd() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	id, _ := this.GetInt("id")
	username := this.GetString("username")
	status, _ := this.GetInt("status")
	password := this.GetString("password")
	roleid, _ := this.GetInt("role_id")
	if id != 0 {
		result := models.RbacUser{Id: id}
		if M.Object().Read(&result) == nil {
			result.Password = password
			result.Status = status
			result.Username = username
			if _, err := M.Object().Update(&result); err == nil {
				tmpResult := models.RbacRoleUser{UserId: id}
				if M.Object().Read(&tmpResult) == nil {
					tmpResult.RoleId = roleid
					M.Object().Update(&tmpResult)
				}
				this.JsonEncode(0, "update success", nil, 0)
			}
			this.JsonEncode(102, "update failed", nil, 0)
		} else {
			this.JsonEncode(102, "update failed", nil, 0)
		}
	}

	result := models.RbacUser{Username: username, Password: password, Status: status}
	uid, err := M.Object().Insert(&result)
	if err != nil {
		this.JsonEncode(101, "insert failed", nil, 0)
	}
	roleuserResult := models.RbacRoleUser{RoleId: roleid, UserId: int(uid)}
	_, err = M.Object().Insert(&roleuserResult)
	this.JsonEncode(0, "success", nil, 0)
}

func (this *MainController) UserDelete() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	id, _ := this.GetInt("id")
	result := models.RbacUser{Id: id}
	_, err := M.Object().Delete(&result)
	if err != nil {
		this.JsonEncode(103, "delete failed", nil, 0)
	}
	//同时删除角色对应关系表
	tmpResult := models.RbacRoleUser{UserId: id}
	M.Object().Delete(&tmpResult)
	this.JsonEncode(0, "delete success", nil, 0)
}

func (this *MainController) AccessListJson() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	roleid, _ := this.GetInt("role_id")
	result := []models.RbacAccess{}
	n, err := M.Object().QueryTable(accessModel).Filter("role_id", roleid).All(&result, "Id", "RoleId", "NodeId", "Level", "Module")
	if err != nil {

	}
	this.JsonEncode(0, "success", result, int(n))
}

func (this *MainController) AccessAdd() {
	_, action := this.GetControllerAndAction()
	this.CheckRbacAll(action)
	roleid, _ := this.GetInt("role_id")
	if roleid == 0 {
		this.JsonEncode(0, "failed", nil, 0)
	}
	str := this.GetString("params")
	data := []string{}
	if len(str) > 0 {
		data = strings.Split(str, ",")
	} else {
		this.JsonEncode(0, "failed", nil, 0)
	}
	//清空某个角色的权限，再重新添加新的
	M.Object().QueryTable(accessModel).Filter("role_id", roleid).Delete()
	result := models.RbacAccess{}
	for _, v := range data {
		tmpStr := strings.Split(v, "_")
		nodeid, _ := strconv.Atoi(tmpStr[0])
		level, _ := strconv.Atoi(tmpStr[1])
		tmpModule := ""
		if level == 1 {
			tmpModule = "项目"
		} else if level == 2 {
			tmpModule = "模块"
		} else {
			tmpModule = "操作"
		}
		result.RoleId = roleid
		result.Level = level
		result.Module = tmpModule
		result.NodeId = nodeid
		M.Object().Insert(&result)
	}
	this.JsonEncode(0, "success", nil, 0)
}
