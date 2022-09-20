package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
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

func (c *CategoryController) Create() {
	c.Layout = "admin_layout.tpl"
	c.TplName = "category/create.tpl"
}

func (c *CategoryController) Post() {
	// 接收数据
	alias := c.GetString("alias")
	name := c.GetString("name")
	sort := c.GetString("sort")
	// 验证数据
	if alias == "" {
		commons.Fail(c.Ctx, "别名不能为空", "", "")
	}
	if name == "" {
		commons.Fail(c.Ctx, "名字不能为空", "", "")
	}
	sortInt, _ := strconv.Atoi(sort)
	// 保存数据
	cat := models.Category{
		Alias: alias,
		Name:  name,
		Sort:  uint(sortInt),
	}
	insert, err := orm.NewOrm().Insert(&cat)
	if err != nil {
		commons.Fail(c.Ctx, "添加失败", nil, "")
	}
	fmt.Println(name, alias)
	// 响应
	commons.Success(c.Ctx, insert, "添加成功", "")
}

func (c *CategoryController) Edit() {
	c.Layout = "admin_layout.tpl"
	c.TplName = "category/edit.tpl"
}
