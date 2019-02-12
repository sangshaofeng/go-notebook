// 注册
package controllers

import (
	"fmt"
	// "encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go-notebook/models"
)

type SignupController struct {
	beego.Controller
}

type User struct {
	Id int
	Username string
	Password string
}

func init() {
	orm.RegisterModel(new(User))
}

func (c *SignupController) Post() {
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	o := orm.NewOrm()
	var user User
	user.Username = username
	user.Password = password
	User.
	id, err := o.Insert(&user)
	fmt.Println(id)
	if err == nil {
		response := models.ResponseBody{
			Code: true,
			Text: "注册成功",
			Data: id,
		}
		c.Data["json"] = &response
		c.ServeJSON()
	}
}
