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
	openingThemes := mtheme.GetThemesByStatus(mtheme.ThemeStatusOpening)
	playingRecordThemes := mrecord.GetThemesByStatus(mrecord.ThemeStatusPlaying)

	//TODO reserve maybe filtered by time
	waitingReserveThemes := mreserve.GetThemesByStatus(mreserve.ThemeStatusWaiting)

	workbenchThemes := make([]*WorkbenchTheme, 0)
	for i := range openingThemes {
		t := openingThemes[i]
		wtheme := new(WorkbenchTheme)
		wtheme.Theme = &t
		for j := range playingRecordThemes {
			r := playingRecordThemes[j]
			if r.Reserve.Theme.Id == wtheme.Theme.Id {
				wtheme.Running = &r
				break
			}
		}
		wtheme.Reserves = make([]*mreserve.Theme, 0)
		for j := range waitingReserveThemes {
			r := waitingReserveThemes[j]
			if r.Theme.Id == wtheme.Theme.Id {
				wtheme.Reserves = append(wtheme.Reserves, &r)
			}
		}
		workbenchThemes = append(workbenchThemes, wtheme)
	}

	c.Data["Themes"] = workbenchThemes
	c.TplName = "dashboard/workbench.html"
}
