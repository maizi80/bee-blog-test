package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	var categorys []models.Category
	o := orm.NewOrm()
	o.QueryTable(new(models.Category)).All(&categorys)
	c.Data["categorys"] = categorys

}
func (c *MainController) SetPaginator(per int, nums int64) *commons.Paginator {
	p := commons.NewPaginator(c.Ctx.Request, per, nums)
	c.Data["paginator"] = p
	return p
}
