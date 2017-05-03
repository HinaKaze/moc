package dashboard

import (
	"github.com/astaxie/beego"
	mrecord "github.com/hinakaze/moc/models/record"
	mreserve "github.com/hinakaze/moc/models/reserve"
	mtheme "github.com/hinakaze/moc/models/theme"
)

type WorkbenchController struct {
	beego.Controller
}

type WorkbenchTheme struct {
	Theme    *mtheme.Theme
	Running  *mrecord.Theme
	Reserves []*mreserve.Theme
}

func (c *WorkbenchController) Get() {
	themes := mtheme.GetAvailableThemes()
	runningRecords := mrecord.GetRunningRecords()
	reserves := mreserve.GetTodayReserves()

	workbenchThemes := make([]*WorkbenchTheme, 0)
	for _, t := range themes {
		wtheme := new(WorkbenchTheme)
		wtheme.Theme = &t
		for _, r := range runningRecords {
			if r.Reserve.Theme.Id == wtheme.Theme.Id {
				wtheme.Running = &r
				break
			}
		}
		wtheme.Reserves = make([]*mreserve.Theme, 0)
		for _, r := range reserves {
			if r.Theme.Id == wtheme.Theme.Id {
				wtheme.Reserves = append(wtheme.Reserves, &r)
			}
		}
		workbenchThemes = append(workbenchThemes, wtheme)
	}

	c.Data["Themes"] = workbenchThemes
	c.TplName = "dashboard/workbench.html"
}
