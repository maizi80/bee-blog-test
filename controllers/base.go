package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	uid := c.GetSession("uid")
	if uid == nil {
		c.Redirect("/login", 302)
	}
}
