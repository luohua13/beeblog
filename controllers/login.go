package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
)

var GlobalSessions *session.Manager
//chrome://settings/content/cookies
func init() {
	sessionconf := &session.ManagerConfig{
		CookieName:"gosessionid", 
		EnableSetCookie: true, 
		Gclifetime:3600,
		Maxlifetime: 3600, 
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory", sessionconf)

	go GlobalSessions.GC()
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	sess, _ := GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	//autoLogin := this.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		
		// Set the UserID if everything is ok
		sess.Set("mySession", "luohua026")
		userID := sess.Get("mySession")
		beego.Debug("======SetSessionId: ",userID,"=========")
	}
	this.Redirect("/",301)
	return
}

func (this *LoginController) Exit() {
	beego.Debug("========Exit==========")
	sess, _ := GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userID := sess.Get("mySession")
	if userID != nil {
		// UserID is set and can be deleted
		sess.Delete("mySession")
	}
	this.Redirect("/",301)
	return
}

func checkAccount(ctx *context.Context) bool {
	sess, _ := GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.ResponseWriter)
	userID := sess.Get("mySession")
	beego.Debug("========CheckSessionId:",userID,"==========")
	if userID == nil {
		// User is logged in already, display another page
		//return false
		return false
	} 
	
	return true
}


