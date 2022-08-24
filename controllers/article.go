package controllers

import (
	"bee-blog/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) Get() {
	id := a.Ctx.Input.Param(":id")
	aid, _ := strconv.ParseUint(id, 0, 8)
	article := models.Article{Id: uint(aid)}
	o := orm.NewOrm()
	if err := o.Read(&article); err != nil {
		a.Abort("404")
	}
	qs := o.QueryTable("article")
	// 上一篇
	var pre models.Article
	qs.Filter("id__lt", uint(aid)).OrderBy("-id").Limit(1).One(&pre)
	fmt.Println(pre)

	// 下一篇
	var next models.Article
	qs.Filter("id__gt", uint(aid)).OrderBy("id").Limit(1).One(&next)
	fmt.Println(next)

	a.Data["article"] = article
	a.Data["Title"] = article.Title
	a.Data["pre"] = pre
	a.Data["next"] = next
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Scripts"] = "a_js.tpl"
	a.LayoutSections["HtmlHead"] = "a_head.tpl"
	a.LayoutSections["Comment"] = "comment.tpl"
	a.Layout = "layout.tpl"
	a.TplName = "article.tpl"
}
