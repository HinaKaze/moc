package reserve

import (
	"time"

	"github.com/astaxie/beego"

	"strconv"

	"github.com/hinakaze/moc/models/record"
	"github.com/hinakaze/moc/models/reserve"
)

type ThemeController struct {
	beego.Controller
}

func (r *ThemeController) Post() {
	teamName := r.GetString("team_name")
	phoneNumber := r.GetString("phone_number")
	memberCount, err := r.GetInt("member_count")
	if err != nil {
		panic(err)
	}
	beginTimeStr := r.GetString("begin_time")

	beginTime, err := time.ParseInLocation("2006-01-02 15:04:05 +0800 CST", beginTimeStr, time.Local)
	if err != nil {
		panic(err)
	}

	newReserveTheme := new(reserve.Theme)
	newReserveTheme.TeamName = teamName
	newReserveTheme.PhoneNumber = phoneNumber
	newReserveTheme.MemberCount = memberCount
	newReserveTheme.BeginTime = beginTime
	reserve.InsertTheme(newReserveTheme)

	r.Redirect("/", 302)
}

func (r *ThemeController) Start() {
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

	reserveTheme, ok := reserve.GetTheme(int64(reserveThemeId))
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
