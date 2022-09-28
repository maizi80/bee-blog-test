package controllers

import (
	"bee-blog/commons"
	"bee-blog/models"
	"bee-blog/services"
	"github.com/beego/beego/v2/client/orm"
	"github.com/spf13/viper"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) Get() {
	var profiles []models.Profile
	orm.NewOrm().QueryTable(new(models.Profile)).All(&profiles)
	m := make(map[string]string)
	for _, v := range profiles {
		m[v.Alias] = v.Content
	}
	for k, _ := range viper.GetStringMap("profile") {
		if _, o := m[k]; !o {
			m[k] = ""
			if k == "avatar" {
				m[k] = viper.GetString("default_avatar")
			}
		}
	}
	c.Data["m"] = m
	c.Layout = "admin_layout.tpl"
	c.TplName = "profile/index.tpl"
}

func (c *ProfileController) Post() {
	f, h, err := c.GetFile("avatar")
	avatar := ""
	if err != nil {
		avatar = viper.GetString("default_avatar")
	} else {
		avatar = "static/upload/" + h.Filename
		defer f.Close()
		c.SaveToFile("avatar", avatar)
	}
	o := orm.NewOrm()
	var num int64
	var profile models.Profile
	for k, v := range viper.GetStringMap("profile") {
		a := c.GetString(k)
		if a == "undefined" {
			continue
		}
		e := o.QueryTable("profile").Filter("alias", k).One(&profile, "id")
		pid := profile.Id

		profile = services.GetProfileByMap(v, models.Profile{})
		profile.Content = a
		if e == nil {
			profile.Id = pid
			if k == "avatar" && err == nil {
				profile.Content = avatar
			}
			if _, e := o.Update(&profile); e != nil {
				commons.Fail(c.Ctx, "更新失败", nil, "")
			}
		} else {
			if k == "avatar" {
				profile.Content = avatar
			}
			if _, e := orm.NewOrm().Insert(&profile); e != nil {
				commons.Fail(c.Ctx, "添加失败", nil, "")
			}
		}
	}

	commons.Success(c.Ctx, num, "操作成功", "")
}
