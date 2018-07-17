package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}
