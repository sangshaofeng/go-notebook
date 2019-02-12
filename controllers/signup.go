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

// 注册接口
func (c *SignupController) Post() {

}
