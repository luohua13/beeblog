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
	
	topics, err := models.GetAllTopics(this.Input().Get("cate"),true)
	
	if err != nil {
		beego.Error(err.Error)
	} else {
		this.Data["Topics"] = topics
	}
	
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	
	this.Data["Categories"] = categories
}
