package models

type RbacNode struct {
	Id     int `orm:"pk"`
	Name   string
	Title  string
	Status int
	Remark string
	Sort   int
	Pid    int
	Level  int
	IsShow int
}
