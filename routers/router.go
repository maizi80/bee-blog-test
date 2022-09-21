package routers

import (
	"bee-blog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article/:id", &controllers.ArticleController{})
	beego.Router("/comment/:aid", &controllers.CommentController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/category/add", &controllers.CategoryController{}, "get:Create")
	beego.Router("/category/:cid", &controllers.CategoryController{}, "get:Edit;put:Put")
}
