package controllers

import (
	"bee-blog/models"
	"github.com/beego/beego/v2/client/orm"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Get() {
	var categorys []models.Category
	orm.NewOrm().QueryTable(new(models.Category)).All(&categorys)
	c.Data["categorys"] = categorys
	c.Layout = "admin_layout.tpl"
	c.TplName = "category.tpl"
}
