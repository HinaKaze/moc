package reserve

import (
	"time"

	"github.com/astaxie/beego"

	"strconv"

	"strings"

	"github.com/hinakaze/moc/models/record"
	"github.com/hinakaze/moc/models/reserve"
	"github.com/hinakaze/moc/models/theme"
)

type ThemeController struct {
	beego.Controller
}

func (r *ThemeController) Post() {
	themeId, err := r.GetInt64("theme_id")
	if err != nil {
		r.Abort(err.Error())
	}
	teamName := r.GetString("team_name")
	phoneNumber := r.GetString("phone_number")
	memberCount, err := r.GetInt("member_count")
	if err != nil {
		r.Abort(err.Error())
	}
	beginTimeStr := strings.TrimSpace(r.GetString("begin_time"))

	beginTime, err := time.ParseInLocation("2006-01-02T15:04:00", beginTimeStr, time.Local)
	if err != nil {
		r.Abort(err.Error())
	}

	newReserveTheme := new(reserve.Theme)
	newReserveTheme.TeamName = teamName
	newReserveTheme.PhoneNumber = phoneNumber
	newReserveTheme.MemberCount = memberCount
	newReserveTheme.BeginTime = beginTime
	newReserveTheme.Theme = new(theme.Theme)
	newReserveTheme.Theme.Id = themeId
	reserve.InsertTheme(newReserveTheme)

	r.Redirect("/", 302)
}

func (r *ThemeController) DoDelete() {
	reserveThemeIdStr := r.Ctx.Input.Param(":id")
	reserveThemeId, err := strconv.ParseInt(reserveThemeIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	reserveTheme, ok := reserve.GetTheme(reserveThemeId)
	if !ok {
		r.Abort("500")
	}
	if reserveTheme.Status != reserve.ThemeStatusWaiting {
		r.Abort("500")
	}

	reserveTheme.Status = reserve.ThemeStatusDeleted
	reserve.UpdateTheme(reserveTheme)

	r.Redirect("/", 302)
}

func (r *ThemeController) DoStart() {
	reserveThemeIdStr := r.Ctx.Input.Param(":id")
	reserveThemeId, err := strconv.ParseInt(reserveThemeIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	payType := r.GetString("pay_type")
	payPrice, err := r.GetFloat("pay_price")
	if err != nil {
		panic(err)
	}

	reserveTheme, ok := reserve.GetTheme(reserveThemeId)
	if !ok {
		r.Redirect("/error", 500)
		return
	}
	if reserveTheme.Status != reserve.ThemeStatusWaiting {
		r.Redirect("/error", 500)
		return
	}

	//update reserve
	reserveTheme.Status = reserve.ThemeStatusConverted
	reserve.UpdateTheme(reserveTheme)

	//create new record
	newRecordTheme := record.CreateTheme(reserveTheme, payType, payPrice)
	record.InsertTheme(newRecordTheme)

	r.Redirect("/", 302)
}
