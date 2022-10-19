package controllers

import (
	"bee-blog/models"
	"bee-blog/services"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type HomeController struct {
	MainController
}

func (c *HomeController) Get() {
	o := orm.NewOrm()

	var articles []models.Article
	o.QueryTable(new(models.Article)).Limit(3).All(&articles)

	var profile []models.Profile
	o.QueryTable(new(models.Profile)).All(&profile)
	p := services.GetMapByProfile(profile)
	fmt.Println(p)
	c.Data["articles"] = articles
	c.Data["p"] = p
	c.Data["Title"] = "首页"
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
