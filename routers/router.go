package routers

import (
	"go-notebook/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
}
