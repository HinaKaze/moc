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
	for i := range themes {
		t := themes[i]
		wtheme := new(WorkbenchTheme)
		wtheme.Theme = &t
		for j := range runningRecords {
			r := runningRecords[j]
			if r.Reserve.Theme.Id == wtheme.Theme.Id {
				wtheme.Running = &r
				break
			}
		}
		wtheme.Reserves = make([]*mreserve.Theme, 0)
		for j := range reserves {
			r := reserves[j]
			if r.Theme.Id == wtheme.Theme.Id {
				wtheme.Reserves = append(wtheme.Reserves, &r)
			}
		}
		workbenchThemes = append(workbenchThemes, wtheme)
	}

	c.Data["Themes"] = workbenchThemes
	c.TplName = "dashboard/workbench.html"
}
