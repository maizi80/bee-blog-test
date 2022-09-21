package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

// TagController operations for Tag
type TagController struct {
	BaseController
}

// URLMapping ...
func (c *TagController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *TagController) Create() {
	c.Layout = "admin_layout.tpl"
	c.TplName = "tag/create.tpl"
}

// Post ...
// @Title Post
// @Description create Tag
func (c *TagController) Post() {
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
	tag := models.Tag{
		Alias: alias,
		Name:  name,
		Sort:  sortInt,
	}
	insert, err := orm.NewOrm().Insert(&tag)
	if err != nil {
		commons.Fail(c.Ctx, "添加失败", nil, "")
	}
	// 响应
	commons.Success(c.Ctx, insert, "添加成功", "")
}

// GetOne ...
// @Title Get One
// @Description get Tag by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tag
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TagController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTagById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Tag
func (c *TagController) GetAll() {
	var tags []models.Tag
	orm.NewOrm().QueryTable(new(models.Tag)).All(&tags)
	c.Data["tags"] = tags
	c.Layout = "admin_layout.tpl"
	c.TplName = "tag/index.tpl"
}

func (c *TagController) Edit() {
	tid := c.Ctx.Input.Param(":tid")
	tidInt, _ := strconv.ParseInt(tid, 0, 64)
	// 验证数据
	if tidInt == 0 {
		commons.Fail(c.Ctx, "ID不能为空", "", "")
	}
	o := orm.NewOrm()
	tag := models.Tag{Id: tidInt}
	err := o.Read(&tag)
	if err == orm.ErrNoRows {
		c.Abort("404")
	}
	c.Data["tag"] = tag
	c.Layout = "admin_layout.tpl"
	c.TplName = "tag/edit.tpl"
}

// Put ...
// @Title Put
// @Description update the Tag
func (c *TagController) Put() {
	tid := c.Ctx.Input.Param(":tid")
	tidInt, _ := strconv.ParseInt(tid, 0, 64)
	alias := c.GetString("alias")
	name := c.GetString("name")
	sort := c.GetString("sort")
	// 验证数据
	if tidInt == 0 {
		commons.Fail(c.Ctx, "ID不能为空", "", "")
	}
	if alias == "" {
		commons.Fail(c.Ctx, "别名不能为空", "", "")
	}
	if name == "" {
		commons.Fail(c.Ctx, "名字不能为空", "", "")
	}
	sortInt, _ := strconv.Atoi(sort)
	// 保存数据
	o := orm.NewOrm()
	err := o.Read(&models.Tag{Id: tidInt})
	if err == orm.ErrNoRows {
		commons.Fail(c.Ctx, "数据错误", nil, "")
	}

	// 保存数据
	tag := models.Tag{
		Id:    tidInt,
		Alias: alias,
		Name:  name,
		Sort:  sortInt,
	}
	num, err := o.Update(&tag)
	if err != nil {
		commons.Fail(c.Ctx, "更新失败", nil, "")
	}
	commons.Success(c.Ctx, num, "更新成功", "")
}

// Delete ...
// @Title Delete
// @Description delete the Tag
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TagController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteTag(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
