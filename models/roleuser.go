package models

type RbacRoleUser struct {
	RoleId int
	UserId int `orm:"pk"`
}
