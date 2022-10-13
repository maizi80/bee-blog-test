package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"bee-blog/util"
	"bee-blog/validations"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
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
		commons.Fail(c.Ctx, "添加失败", "", "")
	}
	commons.Success(c.Ctx, insert, "添加成功", "")
}

func (c *AdminController) Edit() {
	var categorys []models.Category
	var tags []models.Tag
	o := orm.NewOrm()
	o.QueryTable(new(models.Category)).All(&categorys)
	o.QueryTable(new(models.Tag)).All(&tags)

	aid := c.Ctx.Input.Param(":aid")
	aidInt, _ := strconv.Atoi(aid)
	// 验证数据
	if aidInt == 0 {
		commons.Fail(c.Ctx, "ID不能为空", "", "")
	}
	article := models.Article{Id: uint(aidInt)}
	err := o.Read(&article)
	if err == orm.ErrNoRows {
		c.Abort("404")
	}
	c.Data["a"] = article
	c.Data["categorys"] = categorys
	c.Data["tags"] = tags
	c.Data["contains"] = util.ContainsIntInString
	c.Layout = "admin_layout.tpl"
	c.TplName = "admin/edit.tpl"
}

func (c *AdminController) Put() {
	aid := c.Ctx.Input.Param(":aid")
	aidInt, _ := strconv.Atoi(aid)
	cid, _ := c.GetInt("category_id", 0)
	f, h, err := c.GetFile("image")
	image := c.GetString("img")
	if err != nil {
		image = c.GetString("img")
	} else {
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
	o := orm.NewOrm()
	aerr := o.Read(&models.Article{Id: uint(aidInt)})
	if aerr == orm.ErrNoRows {
		commons.Fail(c.Ctx, "数据错误", nil, "")
	}

	// 保存数据
	ao := models.Article{
		Id:           uint(aidInt),
		Title:        a.Title,
		Introduction: a.Introduction,
		Content:      a.Content,
		Image:        a.Image,
		CategoryId:   uint(a.CategoryId),
		Tag:          util.Convert(a.Tag),
	}

	num, err := orm.NewOrm().Update(&ao)
	if err != nil {
		commons.Fail(c.Ctx, "更新失败", nil, "")
	}
	commons.Success(c.Ctx, num, "更新成功", "")
}

func (c *AdminController) Delete() {
	aid := c.Ctx.Input.Param(":aid")
	aidInt, _ := strconv.Atoi(aid)
	// 验证数据
	if aidInt == 0 {
		commons.Fail(c.Ctx, "ID不能为空", "", "")
	}
	a := models.Article{Id: uint(aidInt)}
	o := orm.NewOrm()
	err := o.Read(&a)
	if err == orm.ErrNoRows {
		commons.Fail(c.Ctx, "数据错误", nil, "")
	}
	num, err := o.Delete(&a)
	if err != nil {
		commons.Fail(c.Ctx, "删除失败", nil, "")
	}
	commons.Success(c.Ctx, num, "删除成功", "")
}
