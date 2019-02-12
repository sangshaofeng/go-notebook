package routers

import (
	"go-notebook/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/signup", &controllers.SignupController{})
}
