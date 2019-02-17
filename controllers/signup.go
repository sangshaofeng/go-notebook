// 注册
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	. "go-notebook/models"
	"go-notebook/utils"
)

type SignupController struct {
	beego.Controller
}

// 注册
func (c *SignupController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if !utils.CheckUsername(username) {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "illegal username", }
		c.ServeJSON()
		return
	}
	id, err := AddUser(username, password)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "注册失败",}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "msg": fmt.Sprintf("[%d] ", id) + "注册成功", }
	}
	c.ServeJSON()
}
