package routers

import (
	"github.com/astaxie/beego"
	controller "github.com/hinakaze/moc/controllers"
	"github.com/hinakaze/moc/controllers/dashboard"
	"github.com/hinakaze/moc/controllers/record"
	"github.com/hinakaze/moc/controllers/reserve"
	"github.com/hinakaze/moc/controllers/theme"
)

func init() {
	beego.Router("/", &controller.MainController{})

	beego.Router("/dashboard/workbench", &dashboard.WorkbenchController{})

	beego.Router("/theme/:id/rank", &theme.RankController{}, "get:GetRank")

	beego.Router("/reserve/theme", &reserve.ThemeController{}, "post:Post")
	beego.Router("/reserve/theme/:id/start", &reserve.ThemeController{}, "get:DoStart")
	beego.Router("/reserve/theme/:id/update", &reserve.ThemeController{}, "post:DoUpdate")
	beego.Router("/reserve/theme/:id/delete", &reserve.ThemeController{}, "get:DoDelete")

	beego.Router("/record/theme/:id/finish", &record.ThemeController{}, "get:DoFinish")
	beego.Router("/record/theme/:id/unfinish", &record.ThemeController{}, "get:DoUnfinish")
	beego.Router("/record/theme/:id/tip", &record.ThemeController{}, "get:DoTip")
}
