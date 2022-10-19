package controllers

import (
	"bee-blog/commons"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	uid := c.GetSession("uid")
	if uid == nil {
		//c.Redirect("/login", 302)
	}
}

func (c *BaseController) SetPaginator(per int, nums int64) *commons.Paginator {
	p := commons.NewPaginator(c.Ctx.Request, per, nums)
	c.Data["paginator"] = p
	return p
}
