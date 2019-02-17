package models

import (
	"time"
	"go-notebook/utils"
	"github.com/astaxie/beego/orm"
)
 
type Users struct {
	Id int
	Username string
	Password string
	Salt string
	CreatedAt int64
}

func init() {
	orm.RegisterModel(new(Users))
}

// 添加用户
func AddUser(username string, password string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.Username = username
	user.Salt = utils.RandString(10)
	user.Password = utils.Md5(password + user.Salt)
	user.CreatedAt = time.Now().Unix()
	return o.Insert(user)
}

// 用户名查找用户
func FindUser(username string) (Users, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username: username}
	err := o.Read(&user, "username")
	return user, err
}

// 登录

// 修改密码