package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (e *ErrorController) Get() {
	e.TplName = "error.html"
}
