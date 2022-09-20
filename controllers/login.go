package controllers

import (
	"bee-blog/commons"
	"bee-blog/services"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	uid := c.GetSession("uid")
	fmt.Println(uid)
	if uid != nil {
		c.Redirect("/admin", 302)
	}
	c.TplName = "login.tpl"
}
func (c *LoginController) Post() {
	// 接收数据
	username := c.GetString("username")
	password := c.GetString("password")
	// 验证数据
	if username == "" {
		commons.Fail(c.Ctx, "用户名不能为空", "", "")
	}
	if password == "" {
		commons.Fail(c.Ctx, "密码不能为空", "", "")
	}
	uid := services.CheckUser(username, password)
	if uid == 0 {
		commons.Fail(c.Ctx, "用户或者密码错误", "", "")
	}
	// 保存session
	c.SetSession("uid", uid)
	c.SetSession("username", username)
	fmt.Println(c.GetSession("uid"), c.GetSession("username"))
	// 响应
	commons.Success(c.Ctx, username, "登录成功", "")
}
