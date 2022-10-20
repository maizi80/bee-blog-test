package controllers

import (
	"bee-blog/models"
	"bee-blog/services"
	"github.com/beego/beego/v2/client/orm"
	"math"
	"strconv"
)

type HomeController struct {
	MainController
}

func (c *HomeController) Get() {
	o := orm.NewOrm()
	page := c.Ctx.Input.Param(":page")
	if page == "" {
		page = "1"
	}
	pageInt, _ := strconv.Atoi(page)
	var articles []models.Article
	limit := 5
	offset := (pageInt - 1) * limit
	qs := o.QueryTable(new(models.Article)).Filter("status", 1).Limit(5).OrderBy("-is_top", "sort", "-published_at")
	qs.Limit(limit, offset).All(&articles)

	count, _ := qs.Count()
	co := int(math.Ceil(float64(count) / float64(limit)))

	var recomends []models.Article
	o.QueryTable(new(models.Article)).Filter("status", 1).Filter("is_recommend", 1).OrderBy("-is_top", "sort", "-published_at").All(&recomends)

	var profile []models.Profile
	o.QueryTable(new(models.Profile)).All(&profile)
	p := services.GetMapByProfile(profile)

	c.Data["articles"] = articles
	c.Data["recomends"] = recomends
	c.Data["p"] = p
	c.Data["co"] = co
	c.Data["Title"] = "首页"
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
