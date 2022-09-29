package controllers

import (
	"bee-blog/models"
	"github.com/beego/beego/v2/client/orm"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) Get() {
	var articles []models.Article
	orm.NewOrm().QueryTable(new(models.Article)).All(&articles)

	c.Data["articles"] = articles
	c.Layout = "admin_layout.tpl"
	c.TplName = "admin/index.tpl"
}
