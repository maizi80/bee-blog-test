package main

import (
	_ "bee-blog/inits"
	"bee-blog/models"
	_ "bee-blog/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	o := orm.NewOrm()
	var article models.Article
	article.Title = "title2"
	article.Introduction = "introduction"
	article.Content = "content"
	article.Image = "image"
	article.CategoryId = 1

	id, err := o.Insert(&article)
	if err == nil {
		fmt.Printf("id:%d\n", id)
	}

	a := models.Article{Id: 1}
	e := o.Read(&a)

	if e == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if e == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Printf("id:%d Title:%s\n", a.Id, a.Title)
	}

	beego.Run()
}
