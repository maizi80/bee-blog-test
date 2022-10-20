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
	c.Data["Title"] = "分类-" + category.Name
	category.Name = "分类 “" + category.Name + "” 下"

	c.Data["articles"] = articles
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

	c.Data["Title"] = "标签-" + tag.Name
	tag.Name = "标签 “" + tag.Name + "” 下"
	c.Data["articles"] = articles
	c.Data["co"] = co
	c.Data["count"] = count
	c.Data["list"] = tag
	c.Data["type"] = "tag"
	c.Layout = "layout.tpl"
	c.TplName = "list.tpl"
}

func (c *HomeController) Search() {
	key := c.GetString("s")

	page := c.Ctx.Input.Param(":page")
	if page == "" {
		page = "1"
	}
	pageInt, _ := strconv.Atoi(page)

	limit := 2
	offset := (pageInt - 1) * limit

	o := orm.NewOrm()
	var articles []*models.Article
	qs := o.QueryTable("article").Filter("title__icontains", key).Filter("status", 1)
	qs.Limit(limit, offset).OrderBy("-is_top", "sort", "-published_at").RelatedSel().All(&articles)
	count, _ := qs.Count()
	co := int(math.Ceil(float64(count) / float64(limit)))

	list := make(map[string]string)
	list["Name"] = "包含关键字“" + key + "”的文章"
	list["Id"] = key

	c.Data["Title"] = "搜索-" + key
	c.Data["articles"] = articles

	c.Data["co"] = co
	c.Data["count"] = count
	c.Data["list"] = list
	c.Data["type"] = "search"
	c.Layout = "layout.tpl"
	c.TplName = "list.tpl"
}

func (c *HomeController) Message() {
	aid := "100000"
	pageInt, _ := strconv.Atoi(aid)
	// 评论列表
	commons := services.GetComments(pageInt)
	article := make(map[string]string)
	article["Id"] = aid

	// 获取session
	uid := c.GetSession("uid")
	username := c.GetSession("username")

	c.LayoutSections = make(map[string]string)

	c.Data["Title"] = "留言"
	c.Data["commons"] = commons
	c.Data["article"] = article
	c.Data["uid"] = uid
	c.Data["username"] = username
	c.LayoutSections["Scripts"] = "a_js.tpl"
	c.LayoutSections["Comment"] = "comment.tpl"
	c.Layout = "layout.tpl"
	c.TplName = "message.tpl"
}
func (c *HomeController) About() {
	aid := "200000"
	pageInt, _ := strconv.Atoi(aid)
	// 评论列表
	commons := services.GetComments(pageInt)
	article := make(map[string]string)
	article["Id"] = aid

	// 获取session
	uid := c.GetSession("uid")
	username := c.GetSession("username")

	c.LayoutSections = make(map[string]string)
	c.Data["Title"] = "关于博主"
	c.Data["commons"] = commons
	c.Data["article"] = article
	c.Data["uid"] = uid
	c.Data["username"] = username
	c.LayoutSections["Scripts"] = "a_js.tpl"
	c.LayoutSections["Comment"] = "comment.tpl"
	c.Layout = "layout.tpl"
	c.TplName = "about.tpl"
}
