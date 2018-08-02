package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
	"fmt"
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
			break
		}
		
		err := models.AddCategory(name,true)
		if err != nil {
			fmt.Println("oooooooooooop")
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	}
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsCategory"] = true
	
	
	var err error
	this.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}
	
	this.TplName = "category.html"
}

