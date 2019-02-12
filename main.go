package main

import (
	_ "go-notebook/routers"
	_ "go-notebook/initial"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

