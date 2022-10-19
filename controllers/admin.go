package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"bee-blog/util"
	"bee-blog/validations"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) Get() {
	var articles []*models.Article
	limit := 3

	qs := orm.NewOrm().QueryTable("article").Filter("Category__gt", 0)
	nums, _ := qs.Count()
	pager := c.SetPaginator(limit, nums)
	qs.Limit(limit, pager.Offset()).OrderBy("-created_at").RelatedSel().All(&articles)
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
	fmt.Println(err, h)
	image := ""
	if err == nil {
		image = "static/upload/" + h.Filename
		defer f.Close()
		c.SaveToFile("image", image)
		image = "/" + image
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

	category := models.Category{Id: uint(a.CategoryId)}
	ao := models.Article{
		Title:        a.Title,
		Introduction: a.Introduction,
		Content:      a.Content,
		Image:        a.Image,
		Category:     &category,
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
	fmt.Println(err, h.Filename)
	image := c.GetString("img")
	if err != nil {
		image = c.GetString("img")
	} else {
		image = "static/upload/" + h.Filename
		defer f.Close()
		c.SaveToFile("image", image)
		image = "/" + image
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
	category := models.Category{Id: uint(a.CategoryId)}
	// 保存数据
	ao := models.Article{
		Id:           uint(aidInt),
		Title:        a.Title,
		Introduction: a.Introduction,
		Content:      a.Content,
		Image:        a.Image,
		Category:     &category,
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

func (c *AdminController) ChangeStatus() {
	aid := c.Ctx.Input.Param(":aid")
	aidInt, _ := strconv.Atoi(aid)
	t := c.Ctx.Input.Param(":type")
	status := c.Ctx.Input.Param(":status")
	statusInt, _ := strconv.Atoi(status)
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
	// 保存数据
	switch t {
	case "status":
		a.Status = uint(statusInt)
		a.PublishedAt = time.Now()
	case "top":
		a.IsTop = uint(statusInt)
	case "recommend":
		a.IsRecommend = uint(statusInt)
	}
	num, err := orm.NewOrm().Update(&a)
	if err != nil {
		commons.Fail(c.Ctx, "操作失败", nil, "")
	}
	commons.Success(c.Ctx, num, "操作成功", "")
}
