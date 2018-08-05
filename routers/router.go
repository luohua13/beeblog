package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.AutoRouter(&controllers.LoginController{})

	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.CategoryController{})

	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	os.Mkdir("attachment", os.ModePerm)
	//作为静态文件
	//beego.SetStaticPath("/attachment", "attachment")
	//作为一个单独的控制器
	beego.Router("/attachment/:all", &controllers.AttachControllet{})
}
