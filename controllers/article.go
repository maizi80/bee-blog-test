package controllers

import (
	"bee-blog/models"
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
	if err := orm.NewOrm().Read(&article); err != nil {
		a.Abort("404")
	}
	a.Data["article"] = article
	a.Data["Title"] = article.Title
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Scripts"] = "a_js.tpl"
	a.LayoutSections["HtmlHead"] = "a_head.tpl"
	a.LayoutSections["Comment"] = "comment.tpl"
	a.Layout = "layout.tpl"
	a.TplName = "article.tpl"
}
