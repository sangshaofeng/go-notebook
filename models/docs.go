// 文档model
package models

import (
	// "go-notebook/utils"
	"github.com/astaxie/beego/orm"
	"strconv"
)

// 文档
type Documents struct {
	Id int
	ContentId int
	Title string
	Content string
	Keywords string
	Abstract string
	Author string
}

// 文档目录
type DocContents struct {
	Id int
	Name string
}

func init() {
	orm.RegisterModel(new(Documents), new(DocContents))
}

// 获取全部文档目录
func GetAllDocContent(page int) ([]orm.Params, int64, int, error) {
	o := orm.NewOrm()
	sql := "select * from doc_contents order by created_at desc limit ?,10"
	countSql := "select count(*) as number from doc_contents"
	var maps, maps2 []orm.Params
	num, err := o.Raw(sql, (page - 1) * 10).Values(&maps)
	// 注意下面这一行语句要使用 = 而不是使用 := ,因为 := 需要左侧有一个新的变量名字，
	// 而匿名符号 _ 和 err 都不是新的变量名，所以会报错！
	_, err = o.Raw(countSql).Values(&maps2)
	totalCount, err := strconv.Atoi(maps2[0]["number"].(string))
	var totalPages int
	if (totalCount % 10) == 0 {
		totalPages = totalCount / 10
	} else {
		totalPages = totalCount / 10 + 1
	}
	return maps, num, totalPages, err
}

// 创建文档目录
func CreateDocContent(docName string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	docContent := new(DocContents)
	docContent.Name = docName
	return o.Insert(docContent) 
}

