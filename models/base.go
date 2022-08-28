package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o orm.Ormer
)

func init() {
	//orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/rbac?charset=utf8", 30)
	orm.RegisterModel(new(RbacAccess), new(RbacNode), new(RbacRole), new(RbacRoleUser), new(RbacUser))
	o = orm.NewOrm()
}

type M struct {
}

func (this *M) Object() orm.Ormer {
	return o
}
