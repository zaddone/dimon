package routers

import (
	"github.com/astaxie/beego"
	"github.com/zaddone/dimon/controllers/admin"
	"github.com/zaddone/dimon/controllers/blog"
)

func init() {

	beego.Router("/", &blog.MainController{}, "*:Index")
	beego.Router("/404.html", &blog.MainController{}, "*:Go404")
	beego.Router("/index:page:int.html", &blog.MainController{}, "*:Index")

	beego.Router("/article/:id:int", &blog.MainController{}, "*:Show")      //ID访问
	beego.Router("/article/:urlname(.+)", &blog.MainController{}, "*:Show") //别名访问文章

	beego.Router("/category/:name(.+?)", &blog.MainController{}, "*:Category")
	beego.Router("/category/:name(.+?)/page/:page:int", &blog.MainController{}, "*:Category")

	beego.Router("/life:page:int.html", &blog.MainController{}, "*:BlogList")
	beego.Router("/life.html", &blog.MainController{}, "*:BlogList")

	beego.Router("/mood.html", &blog.MainController{}, "*:Mood")
	beego.Router("/mood:page:int.html", &blog.MainController{}, "*:Mood")

	beego.Router("/photo.html", &blog.MainController{}, "*:Photo")
	beego.Router("/photo:page:int.html", &blog.MainController{}, "*:Photo")

	beego.Router("/album.html", &blog.MainController{}, "*:Album")
	beego.Router("/album:page:int.html", &blog.MainController{}, "*:Album")

	beego.Router("/book.html", &blog.MainController{}, "*:Book")
	beego.Router("/:urlname(.+)", &blog.MainController{}, "*:Show") //别名访问

	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")
	beego.Router("/admin/system/setting", &admin.SystemController{}, "*:Setting")

	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article/batch", &admin.ArticleController{}, "*:Batch")
	beego.Router("/admin/article/upload", &admin.ArticleController{}, "*:Upload")
	beego.Router("/admin/tag", &admin.TagController{}, "*:Index")

	beego.Router("/admin/mood/add", &admin.MoodController{}, "*:Add")
	beego.Router("/admin/mood/list", &admin.MoodController{}, "*:List")
	beego.Router("/admin/mood/delete", &admin.MoodController{}, "*:Delete")

	beego.Router("/admin/album/add", &admin.AlbumController{}, "*:Add")
	beego.Router("/admin/album/list", &admin.AlbumController{}, "*:List")
	beego.Router("/admin/album/edit", &admin.AlbumController{}, "*:Edit")
	beego.Router("/admin/album/delete", &admin.AlbumController{}, "*:Delete")

	beego.Router("/admin/photo/list", &admin.PhotoController{}, "*:List")
	beego.Router("/admin/photo/cover", &admin.PhotoController{}, "*:Cover")
	beego.Router("/admin/photo/delete", &admin.PhotoController{}, "*:Delete")
	beego.Router("/admin/photo/uploadphoto", &admin.PhotoController{}, "*:UploadPhoto")

	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

	beego.Router("/admin/link/list", &admin.LinkController{}, "*:List")
	beego.Router("/admin/link/add", &admin.LinkController{}, "*:Add")
	beego.Router("/admin/link/edit", &admin.LinkController{}, "*:Edit")
	beego.Router("/admin/link/delete", &admin.LinkController{}, "*:Delete")
}
