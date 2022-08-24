package controllers

import (
	"bee-blog/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var articles []models.Article
	o := orm.NewOrm().QueryTable(new(models.Article))
	o.Limit(3).All(&articles)
	c.Data["articles"] = articles
	c.Data["Title"] = "首页"
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
