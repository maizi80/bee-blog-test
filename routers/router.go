package routers

import (
	"bee-blog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/article/home/page/:page", &controllers.HomeController{}, "get:Get")
	beego.Router("/article/:id", &controllers.ArticleController{})
	beego.Router("/comment/:aid", &controllers.CommentController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/category/add", &controllers.CategoryController{}, "get:Create")
	beego.Router("/category/:cid", &controllers.CategoryController{}, "get:Edit;put:Put;delete:Delete")
	beego.Router("/tag", &controllers.TagController{}, "get:GetAll;post:Post")
	beego.Router("/tag/add", &controllers.TagController{}, "get:Create")
	beego.Router("/tag/:tid", &controllers.TagController{}, "get:Edit;put:Put;delete:Delete")
	beego.Router("/profile", &controllers.ProfileController{})
	beego.Router("/admin/add", &controllers.AdminController{}, "get:Create")
	beego.Router("/admin/:aid", &controllers.AdminController{}, "get:Edit;put:Put;delete:Delete")
	beego.Router("/admin/:aid/:type/:status", &controllers.AdminController{}, "put:ChangeStatus")
	beego.Router("/list/category/:cid", &controllers.HomeController{}, "get:Category")
	beego.Router("/article/category/page/:cid/:page", &controllers.HomeController{}, "get:Category")
	beego.Router("/list/tag", &controllers.HomeController{}, "get:Tag")
	beego.Router("/list/tag/:tid", &controllers.HomeController{}, "get:TagList")
	beego.Router("/article/tag/page/:tid/:page", &controllers.HomeController{}, "get:Tag")
	beego.Router("/search", &controllers.HomeController{}, "post:Search")
	beego.Router("/article/search/page/:key/:page", &controllers.HomeController{}, "get:Search")
	beego.Router("/message", &controllers.HomeController{}, "get:Message")
	beego.Router("/about", &controllers.HomeController{}, "get:About")
	beego.Router("/archive", &controllers.HomeController{}, "get:Archive")

}
