package models

import (
	"github.com/astaxie/beego/orm"
)
 
type Users struct {
	Id int
	Username string
	Password string
	Salt string
}

func init() {
	
}

func AddUser(username string, password string) (int64, error) {
	o := orm.NewOrm()
}