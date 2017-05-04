package routers

import (
	"github.com/astaxie/beego"
	controller "github.com/hinakaze/moc/controllers"
	"github.com/hinakaze/moc/controllers/dashboard"
	"github.com/hinakaze/moc/controllers/record"
	"github.com/hinakaze/moc/controllers/reserve"
)

func init() {
	beego.Router("/", &controller.MainController{})
	beego.Router("/error", &controller.ErrorController{})

	beego.Router("/dashboard/workbench", &dashboard.WorkbenchController{})

	beego.Router("/reserve/theme", &reserve.ThemeController{}, "post:Post")
	beego.Router("/reserve/theme/:id/start", &reserve.ThemeController{}, "*:Start")

	beego.Router("/record/theme/:id/finish", &record.ThemeController{}, "put:Finish")
	beego.Router("/record/theme/:id/unfinish", &record.ThemeController{}, "put:Unfinish")
	beego.Router("/record/theme/:id/tip", &record.ThemeController{}, "put:Tip")
}
