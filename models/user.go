package models

type RbacUser struct {
	Id        int `orm:"pk"`
	Username  string
	Password  string
	Status    int
	Loginip   string
	Logintime string
}
