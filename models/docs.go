// 文档model
package models

import (
	// "go-notebook/utils"
	"github.com/astaxie/beego/orm"
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
func GetAllDocContent(page int) ([]orm.Params, int64, int64, error) {
	o := orm.NewOrm()
	sql := "select * from doc_contents order by created_at desc limit ?,10"
	countSql := "select count(*) from doc_contents"
	var maps, maps2 []orm.Params
	num, err := o.Raw(sql, (page - 1) * 10).Values(&maps)
	count, err := o.Raw(countSql).Values(&maps2)
	var totalPages int64
	if count % 10 == 0 {
		totalPages = count % 10
	} else {
		totalPages = count % 10 + 1
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

