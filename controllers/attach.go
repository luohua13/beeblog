package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type AttachControllet struct {
	beego.Controller
}

func (this *AttachControllet) Get() {
	filepath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}

	f, Err := os.Open(filepath)
	if Err != nil {
		this.Ctx.WriteString(Err.Error())
	}
	defer f.Close()

	_, err = io.Copy(this.Ctx.ResponseWriter, f)
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
}
