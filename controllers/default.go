package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	str := `
	short url (Redis version) :
	1. http://localhost:8080/url/shorturl/?longurl=http://google.com. 
	2. http://localhost:8080/url/longurl/?shorturl=P84OE7
	`
	this.Ctx.Output.Body([]byte(str))
}
