package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"bee-blog/util"
	"bee-blog/validations"
	"fmt"
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

func (c *AdminController) Create() {
	var categorys []models.Category
	var tags []models.Tag
	o := orm.NewOrm()
	o.QueryTable(new(models.Category)).All(&categorys)
	o.QueryTable(new(models.Tag)).All(&tags)
	c.Data["categorys"] = categorys
	c.Data["tags"] = tags
	c.Layout = "admin_layout.tpl"
	c.TplName = "admin/create.tpl"
}

func (c *AdminController) Post() {
	cid, _ := c.GetInt("category_id", 0)
	f, h, err := c.GetFile("image")
	image := ""
	if err == nil {
		image = "/static/upload/" + h.Filename
		defer f.Close()
		c.SaveToFile("image", image)
	}
	a := validations.ArticleValidate{
		Title:        c.GetString("title"),
		CategoryId:   cid,
		Introduction: c.GetString("introduction"),
		Content:      c.GetString("content"),
		Image:        image,
		Tag:          c.GetStrings("tag"),
	}
	commons.Valid(c.Ctx, &a)

	ao := models.Article{
		Title:        a.Title,
		Introduction: a.Introduction,
		Content:      a.Content,
		Image:        a.Image,
		CategoryId:   uint(a.CategoryId),
		Tag:          util.Convert(a.Tag),
	}
	insert, err := orm.NewOrm().Insert(&ao)
	if err != nil {
		fmt.Printf(err.Error())
		commons.Fail(c.Ctx, "添加失败", "", "")
	}
	fmt.Println(a, insert)
	commons.Success(c.Ctx, insert, "添加成功", "")
}
