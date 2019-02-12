package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:shaofeng1022@tcp(127.0.0.1:3306)/go_notebook?charset=utf8", 30)
}