package models

type RbacRole struct {
	Id       int           `orm:"pk"`
	Name     string
	Pid      int
	Status   int
	Remark   string
}
