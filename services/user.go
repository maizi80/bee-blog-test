package services

import (
	"bee-blog/models"
	"github.com/beego/beego/v2/client/orm"
)

type userService struct {
}

func CheckUser(username string, password string) uint {
	var user models.User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs = qs.Filter("name", username)
	if password != "" {
		qs = qs.Filter("password", password)
	}
	r := qs.One(&user, "Id")
	if r != nil {
		return 0
	}
	return user.Id
}
