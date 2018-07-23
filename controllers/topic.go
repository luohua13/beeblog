package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	//"path"
	//"strings"
)

type TopicController struct {
	beego.Controller
}

//文章列表
func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	
	topics, err := models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}

//增加修改文章，保存到数据库中
func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	
	var err error
	err = models.AddTopic(title,content)
	
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic",301)
}

//增加文章
func (this *TopicController) Add() {
	this.TplName = "topic_add.html"

	//this.Data["CategoryList"], _ = models.GetAllCategory()
}

//修改文章
// func (this *TopicController) Modify() {
	// this.TplNames = "topic_modify.html"

	// tid := this.Input().Get("tid")
	// topic, err := models.GetTopic(tid)
	// if err != nil {
		// beego.Error(err)
		// this.Redirect("/", 302)
		// return
	// }

	// this.Data["Topic"] = topic
	// this.Data["Tid"] = tid

	// this.Data["CategoryList"], _ = models.GetAllCategory()
// }

//文章详情
// func (this *TopicController) View() {
	// this.TplNames = "topic_view.html"

	// topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	// if err != nil {
		// beego.Error(err)
		// this.Redirect("/", 302)
		// return
	// }

	// this.Data["Topic"] = topic
	// this.Data["Labels"] = strings.Split(topic.Labels, " ")
	// this.Data["Tid"] = this.Ctx.Input.Param("0")

	// replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	// if err != nil {
		// beego.Error(err)
		// return
	// }

	// this.Data["Replies"] = replies
	// this.Data["IsLogin"] = checkAccount(this.Ctx)
// }

//删除文章
// func (this *TopicController) Delete() {
	// if !checkAccount(this.Ctx) {
		// this.Redirect("/login", 302)
		// return
	// }

	// err := models.DeleteTopic(this.Ctx.Input.Param("0"))
	// if err != nil {
		// beego.Error(err)
	// }
	// this.Redirect("/topic", 302)

// }