package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	c.Data["json"] = "astaxie@gmail.com"
	c.ServeJSON()
}
