package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"log"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/51hsw?charset=utf8")
	orm.RegisterModel(new(Links))
}

type Links struct {
	Id       int `orm:"pk"`
	Webtitle string
}

func main() {
	link := new(Links)
	o := orm.NewOrm()
	links := []Links{}
	o.QueryTable(link).All(&links)
	log.Println(links)
}
