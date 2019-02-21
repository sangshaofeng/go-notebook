package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	. "go-notebook/models"
)

// 文档目录
type DocContentController struct {
	beego.Controller
}

// 获取文档目录
func (c *DocContentController) Get() {
	user := c.GetSession("username")
	if user == nil {
		c.Data["json"] = map[string]interface{}{"code": 1001, "msg": "未登录",}
		c.ServeJSON()
		return
	}
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	docs, count, totalPages, err := GetAllDocContent(page)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 0, 
			"msg": "请求失败", 
			"data": err,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 1, 
			"msg": "请求成功",
			"data": docs,
			"count": count,
			"totalPages": totalPages, 
		}
	}
	c.ServeJSON()
}

// 创建文档目录
func (c *DocContentController) Post() {
	user := c.GetSession("username")
	if user == nil {
		c.Data["json"] = map[string]interface{}{"code": 1001, "msg": "未登录",}
		c.ServeJSON()
		return
	}
	docName := c.GetString("name")
	id, err := CreateDocContent(docName)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 0, 
			"msg": "请求失败", 
			"data": err,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 1, 
			"msg": "创建成功",
			"data": map[string]int64{"id": id,},
		}
	}
	c.ServeJSON()
}