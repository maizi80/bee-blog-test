package controllers

import (
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

// Post ...
// @Title Post
// @Description create Tag
// @Param	body		body 	models.Tag	true		"body for Tag content"
// @Success 201 {int} models.Tag
// @Failure 403 body is empty
// @router / [post]
func (c *TagController) Post() {
	var v models.Tag
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddTag(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
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
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Tag
// @Failure 403
// @router / [get]
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
