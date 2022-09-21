package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"encoding/json"
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

// Put ...
// @Title Put
// @Description update the Tag
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Tag	true		"body for Tag content"
// @Success 200 {object} models.Tag
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TagController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Tag{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateTagById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
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
