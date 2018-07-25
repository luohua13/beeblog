package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	
	topics, err := models.GetAllTopics(true)
	
	if err != nil {
		beego.Error(err.Error)
	} else {
		this.Data["Topics"] = topics
	}
}
