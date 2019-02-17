package controllers

import (
	"github.com/astaxie/beego"
	. "go-notebook/models"
	"go-notebook/utils"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "invalid request"}
	}
	user, err := FindUser(username)
	if err != nil {
		fmt.Println(username, password)
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "用户不存在"}
	} else {
		pwd := utils.Md5(password + user.Salt)
		if pwd == user.Password {
			c.SetSession("username", username)
			c.Data["json"] = map[string]interface{}{"code": 1, "msg": "用户["+ user.Username +"]登录成功"}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "msg": "密码错误"}
		}
	}
	c.ServeJSON()
}
