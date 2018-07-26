package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
	//"github.com/astaxie/beego/context"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tis := this.Input().Get("pid")
	err := models.AddReply(tid,this.Input().Get("nickname"),
		this.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	}
	
	this.Redirct("/topic/view"+tid,302)
}

