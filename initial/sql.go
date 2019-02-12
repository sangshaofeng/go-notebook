package initial

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user     := "root"
	password := "shaofeng1022"
	host     := "localhost"
	port     := 3306
	dbname   := "go_notebook"

	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user,password, host, port, dbname))
}