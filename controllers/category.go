package controllers

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break;
		}
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break;
		}
	}
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
}
