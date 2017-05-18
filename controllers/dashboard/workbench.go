package dashboard

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/hinakaze/moc/common"
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
	Reserves []WorkbenchReserve
}

type WorkbenchReserve struct {
	TimeRange mtheme.TimeRange
	Reserve   *mreserve.Theme
}

func (c *WorkbenchController) Get() {
	openingThemes := mtheme.GetThemesByStatus(mtheme.ThemeStatusOpening)
	playingRecordThemes := mrecord.GetThemesByStatus(mrecord.ThemeStatusPlaying)

	//TODO reserve maybe filtered by time
	todayReserveThemes := mreserve.GetOneDayThemes(time.Now())

	workbenchThemes := make([]*WorkbenchTheme, 0)
	for i := range openingThemes {
		t := openingThemes[i]
		wtheme := new(WorkbenchTheme)
		wtheme.Theme = &t
		for j := range playingRecordThemes {
			r := playingRecordThemes[j]
			if r.Reserve.Theme.ID == wtheme.Theme.ID {
				wtheme.Running = &r
				break
			}
		}
		wtheme.Reserves = make([]WorkbenchReserve, 0)

	LOOP_TIMERANGE:
		for j := range wtheme.Theme.TimeRange {
			timeRange := wtheme.Theme.TimeRange[j]
			workbenchReserve := WorkbenchReserve{}
			workbenchReserve.TimeRange = *timeRange

			wtheme.Reserves = append(wtheme.Reserves, workbenchReserve)
			for z := range todayReserveThemes {
				r := todayReserveThemes[z]
				if timeRange.From.Before(r.BeginTime) && timeRange.To.After(r.BeginTime) {
					workbenchReserve.Reserve = &r
					continue LOOP_TIMERANGE
				}
			}
		}
		workbenchThemes = append(workbenchThemes, wtheme)
	}

	c.Data["Themes"] = workbenchThemes
	c.TplName = "dashboard/workbench.html"
}

func (r *WorkbenchController) GetReserveHistory() {
	fromTimeStr := r.GetString("from")
	toTimeStr := r.GetString("to")

	fromTime, err := common.ParseTime(fromTimeStr)
	if err != nil {
		r.Abort(err.Error())
	}
	toTime, err := common.ParseTime(toTimeStr)
	if err != nil {
		r.Abort(err.Error())
	}

	themes := mreserve.GetHistoryThemes(fromTime, toTime)

	r.Data["Themes"] = themes
	r.TplName = "dashboard/history.html"
}
