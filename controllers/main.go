package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/moc/models/theme"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	openingThemes := theme.GetThemesByStatus(theme.ThemeStatusOpening)

	c.Data["Themes"] = openingThemes

	c.TplName = "main.html"
}
