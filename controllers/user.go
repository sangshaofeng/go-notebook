package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	. "go-notebook/models"
	"go-notebook/utils"
)

// 注册
type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if !utils.CheckUsername(username) {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "用户名格式错误", }
		c.ServeJSON()
		return
	}
	id, err := AddUser(username, password)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "注册失败", "data": err}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "msg": fmt.Sprintf("[%d] ", id) + "注册成功", }
	}
	c.ServeJSON()
}

// 登录
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

// 修改资料：上传头像，修改昵称
type EditUserInfoController struct {
	beego.Controller
}

func (c *EditUserInfoController) Post() {
	
}