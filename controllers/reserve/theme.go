package reserve

import (
	"github.com/astaxie/beego"

	"strconv"

	"strings"

	"fmt"

	"github.com/hinakaze/moc/common"
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

	beginTime, err := common.ParseTime(beginTimeStr)
	if err != nil {
		r.Abort(err.Error())
	}

	payType := r.GetString("pay_type")
	payPrice, err := r.GetFloat("pay_price")
	if err != nil {
		if payType == "未支付" {
			payPrice = 0.0
		} else {
			r.Abort(err.Error())
		}
	}

	newReserveTheme := new(reserve.Theme)
	newReserveTheme.TeamName = teamName
	newReserveTheme.PhoneNumber = phoneNumber
	newReserveTheme.MemberCount = memberCount
	newReserveTheme.BeginTime = beginTime
	newReserveTheme.Theme = new(theme.Theme)
	newReserveTheme.Theme.Id = themeId
	newReserveTheme.PayType = payType
	newReserveTheme.PayPrice = payPrice
	reserve.InsertTheme(newReserveTheme)

	common.HandleSuccess(&r.Controller)
}

func (r *ThemeController) DoUpdate() {
	reserveThemeIdStr := r.Ctx.Input.Param(":id")
	reserveThemeId, err := strconv.ParseInt(reserveThemeIdStr, 10, 64)
	if err != nil {
		r.Abort(err.Error())
	}
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

	beginTime, err := common.ParseTime(beginTimeStr)
	if err != nil {
		r.Abort(err.Error())
	}

	reserveTheme, ok := reserve.GetTheme(reserveThemeId)
	if !ok {
		r.Abort(fmt.Sprintf("Can't find reserve theme [%d]", reserveThemeId))
	}

	payType := r.GetString("pay_type")
	payPrice, err := r.GetFloat("pay_price")
	if err != nil {
		if payType == "未支付" {
			payPrice = 0.0
		} else {
			r.Abort(err.Error())
		}
	}

	reserveTheme.TeamName = teamName
	reserveTheme.PhoneNumber = phoneNumber
	reserveTheme.MemberCount = memberCount
	reserveTheme.BeginTime = beginTime
	reserveTheme.Theme.Id = themeId
	reserveTheme.PayType = payType
	reserveTheme.PayPrice = payPrice

	reserve.UpdateTheme(reserveTheme)

	common.HandleSuccess(&r.Controller)
}

func (r *ThemeController) DoDelete() {
	reserveThemeIdStr := r.Ctx.Input.Param(":id")
	reserveThemeId, err := strconv.ParseInt(reserveThemeIdStr, 10, 64)
	if err != nil {
		r.Abort(err.Error())
	}

	reserveTheme, ok := reserve.GetTheme(reserveThemeId)
	if !ok {
		r.Abort(fmt.Sprintf("Unkown reserve theme [%d]", reserveThemeId))
	}
	if reserveTheme.Status != reserve.ThemeStatusWaiting {
		r.Abort(fmt.Sprintf("Reserve theme [%d] is not in waiting", reserveThemeId))
	}

	reserveTheme.Status = reserve.ThemeStatusDeleted
	reserve.UpdateTheme(reserveTheme)

	common.HandleSuccess(&r.Controller)
}

func (r *ThemeController) DoStart() {
	reserveThemeIdStr := r.Ctx.Input.Param(":id")
	reserveThemeId, err := strconv.ParseInt(reserveThemeIdStr, 10, 64)
	if err != nil {
		r.Abort(err.Error())
	}

	reserveTheme, ok := reserve.GetTheme(reserveThemeId)
	if !ok {
		r.Abort(fmt.Sprintf("Unkown reserve theme [%d]", reserveThemeId))
	}

	if reserveTheme.PayType == "未支付" {
		r.Abort(fmt.Sprintf("Reserve theme [%d] not payed", reserveThemeId))
	}
	if reserveTheme.Status != reserve.ThemeStatusWaiting {
		r.Abort(fmt.Sprintf("Reserve theme [%d] is not in waiting", reserveThemeId))
	}

	//update reserve
	reserveTheme.Status = reserve.ThemeStatusStarted
	reserve.UpdateTheme(reserveTheme)

	//create new record
	newRecordTheme := record.CreateTheme(reserveTheme, reserveTheme.PayType, reserveTheme.PayPrice)
	record.InsertTheme(newRecordTheme)

	common.HandleSuccess(&r.Controller)
}
