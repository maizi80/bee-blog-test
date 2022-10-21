package controllers

import (
	"bee-blog/models"
	"bee-blog/services"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) Get() {
	id := a.Ctx.Input.Param(":id")
	aid, _ := strconv.ParseUint(id, 0, 8)
	article := models.Article{Id: uint(aid)}
	o := orm.NewOrm()
	// 流量量+1
	o.QueryTable("article").Filter("id", uint(aid)).Update(orm.Params{
		"view_count": orm.ColValue(orm.ColAdd, 1),
	})

	if err := o.Read(&article); err != nil {
		a.Abort("404")
	}
	qs := o.QueryTable("article")
	// 上一篇
	var pre models.Article
	qs.Filter("id__lt", uint(aid)).OrderBy("-id").Limit(1).One(&pre)

	// 下一篇
	var next models.Article
	qs.Filter("id__gt", uint(aid)).OrderBy("id").Limit(1).One(&next)

	var user models.User
	// 获取session
	uid := a.GetSession("uid")
	if uid != nil {
		uidi := uid.(int64)
		//cidInt, _ := strconv.Atoi(uidS)
		user = models.User{Id: uint(uidi)}
		o.Read(&user)
	}

	// 获取博主信息
	var profile []models.Profile
	o.QueryTable(new(models.Profile)).All(&profile)
	p := services.GetMapByProfile(profile)

	// 评论列表
	commons := services.GetComments(int(aid))

	// 标签
	var tag []models.Tag
	o.QueryTable("tag").Filter("id__in", strings.Split(article.Tag, ",")).All(&tag)

	a.Data["article"] = article
	a.Data["Title"] = article.Title
	a.Data["pre"] = pre
	a.Data["next"] = next
	a.Data["commons"] = commons
	a.Data["user"] = user
	a.Data["p"] = p
	a.Data["tag"] = tag
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Scripts"] = "a_js.tpl"
	a.LayoutSections["HtmlHead"] = "a_head.tpl"
	a.LayoutSections["Comment"] = "comment.tpl"
	a.Layout = "layout.tpl"
	a.TplName = "article.tpl"
}
