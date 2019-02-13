package models

import (
	"go-notebook/utils"
	"github.com/astaxie/beego/orm"
)
 
type User struct {
	Id int
	Username string
	Password string
	Salt string
}

func init() {
	orm.RegisterModel(new(User))
}

// 添加用户
func AddUser(username string, password string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Username = username
	user.Salt = utils.RandString(10)
	user.Password = utils.Md5(passord + user.Salt)
	return o.Insert(user)
}

// 用户名查找用户
func FindUser(username strng) (Users, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username: username}
	err := o.Read(&user, "username")
	return user, err
}