package routers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/moc/controllers"
	"github.com/hinakaze/moc/controllers/dashboard"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dashboard/workbench", &dashboard.WorkbenchController{})
}
