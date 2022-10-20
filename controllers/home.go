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

func (c *HomeController) Category() {
	cid := c.Ctx.Input.Param(":cid")
	cidInt, _ := strconv.Atoi(cid)
	page := c.Ctx.Input.Param(":page")
	if page == "" {
		page = "1"
	}
	pageInt, _ := strconv.Atoi(page)

	limit := 2
	offset := (pageInt - 1) * limit

	o := orm.NewOrm()
	var articles []*models.Article
	qs := o.QueryTable("article").Filter("Category__exact", cidInt).Filter("status", 1)
	qs.Limit(limit, offset).OrderBy("-is_top", "sort", "-published_at").RelatedSel().All(&articles)
	count, _ := qs.Count()
	co := int(math.Ceil(float64(count) / float64(limit)))

	var category models.Category
	o.QueryTable("category").Filter("Id", cidInt).One(&category, "Id", "Name")

	c.Data["articles"] = articles
	c.Data["Title"] = "分类-" + category.Name
	c.Data["co"] = co
	c.Data["count"] = count
	c.Data["list"] = category
	c.Data["type"] = "category"
	c.Layout = "layout.tpl"
	c.TplName = "list.tpl"
}

func (c *HomeController) Tag() {
	var tag []models.Tag
	orm.NewOrm().QueryTable("tag").All(&tag)

	c.Data["tag"] = tag
	c.Data["Title"] = "标签云"
	c.Layout = "layout.tpl"
	c.TplName = "tag.tpl"
}

func (c *HomeController) TagList() {
	tid := c.Ctx.Input.Param(":tid")
	tidInt, _ := strconv.Atoi(tid)
	page := c.Ctx.Input.Param(":page")
	if page == "" {
		page = "1"
	}
	pageInt, _ := strconv.Atoi(page)

	limit := 2
	offset := (pageInt - 1) * limit

	o := orm.NewOrm()
	var articles []*models.Article
	qs := o.QueryTable("article").Filter("tag__icontains", tidInt).Filter("status", 1)
	qs.Limit(limit, offset).OrderBy("-is_top", "sort", "-published_at").RelatedSel().All(&articles)
	count, _ := qs.Count()
	co := int(math.Ceil(float64(count) / float64(limit)))

	var tag models.Tag
	o.QueryTable("tag").Filter("Id", tidInt).One(&tag, "Id", "Name")

	c.Data["articles"] = articles
	c.Data["Title"] = "标签-" + tag.Name
	c.Data["co"] = co
	c.Data["count"] = count
	c.Data["list"] = tag
	c.Data["type"] = "tag"
	c.Layout = "layout.tpl"
	c.TplName = "list.tpl"
}
