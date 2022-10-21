package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"bee-blog/util"
	"fmt"
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
	c.TplName = "category/index.tpl"
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
	sortInt, _ := util.ToUInt(sort)
	// 保存数据
	cat := models.Category{
		Alias: alias,
		Name:  name,
		Sort:  sortInt,
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
	cid := c.Ctx.Input.Param(":cid")
	cidInt, _ := util.ToUInt(cid)
	// 验证数据
	if cidInt == 0 {
		commons.Fail(c.Ctx, "ID不能为空", "", "")
	}
	o := orm.NewOrm()
	category := models.Category{Id: cidInt}
	err := o.Read(&category)
	if err == orm.ErrNoRows {
		c.Abort("404")
	}
	c.Data["category"] = category
	c.Layout = "admin_layout.tpl"
	c.TplName = "category/edit.tpl"
}

func (c *CategoryController) Put() {
	cid := c.Ctx.Input.Param(":cid")
	cidInt, _ := util.ToUInt(cid)
	alias := c.GetString("alias")
	name := c.GetString("name")
	sort := c.GetString("sort")
	// 验证数据
	if cidInt == 0 {
		commons.Fail(c.Ctx, "ID不能为空", "", "")
	}
	// 验证数据
	if alias == "" {
		commons.Fail(c.Ctx, "别名不能为空", "", "")
	}
	if name == "" {
		commons.Fail(c.Ctx, "名字不能为空", "", "")
	}
	sortInt, _ := util.ToUInt(sort)
	o := orm.NewOrm()
	err := o.Read(&models.Category{Id: cidInt})
	if err == orm.ErrNoRows {
		commons.Fail(c.Ctx, "数据错误", nil, "")
	}

	// 保存数据
	cat := models.Category{
		Id:    cidInt,
		Alias: alias,
		Name:  name,
		Sort:  sortInt,
	}
	num, err := o.Update(&cat)
	if err != nil {
		commons.Fail(c.Ctx, "更新失败", nil, "")
	}
	commons.Success(c.Ctx, num, "更新成功", "")
}

func (c *CategoryController) Delete() {
	cid := c.Ctx.Input.Param(":cid")
	cidInt, _ := util.ToUInt(cid)
	// 验证数据
	if cidInt == 0 {
		commons.Fail(c.Ctx, "ID不能为空", "", "")
	}
	cat := models.Category{Id: cidInt}
	o := orm.NewOrm()
	err := o.Read(&cat)
	if err == orm.ErrNoRows {
		commons.Fail(c.Ctx, "数据错误", nil, "")
	}
	num, err := o.Delete(&cat)
	if err != nil {
		commons.Fail(c.Ctx, "删除失败", nil, "")
	}
	commons.Success(c.Ctx, num, "删除成功", "")
}
