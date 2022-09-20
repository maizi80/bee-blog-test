package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id        uint      `orm:"column(id);auto;size(11)" description:"表ID"`
	Name      string    `orm:"column(name);size(30);" description:"用户名"`
	Password  string    `orm:"column(password);size(30);" description:"用户名"`
	Email     string    `orm:"column(email);size(30);" description:"邮箱"`
	Code      string    `orm:"column(code);size(10);" description:"验证码"`
	Avatar    string    `orm:"column(avatar);size(255)" description:"头像"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now" description:"添加时间"`
}

func init() {
	orm.RegisterModel(new(User))
}
