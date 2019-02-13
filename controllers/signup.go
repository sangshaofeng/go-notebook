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

// 注册接口
func (c *SignupController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if !utils.CheckUsername(username) {
		c.Data["json"] = map[string]interface{}{"result": false, "msg": "illegal username", "refer": "/"}
		c.ServeJSON()
		return
	}
	id, err := AddUser(username, password)
	fmt.Println(id, err)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"result": false, "msg": "registor failed", "refer": "/"}
	} else {
		c.Data["json"] = map[string]interface{}{"result": true, "msg": fmt.Sprintf("[%d] ", id) + "registor success", "refer": "/"}
	}
	c.ServeJSON()
}
