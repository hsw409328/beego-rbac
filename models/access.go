package models

type RbacAccess struct {
	Id     int `orm:"pk"`
	RoleId int
	NodeId int
	Level  int
	Module string
}
