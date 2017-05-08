package record

import (
	"strconv"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/hinakaze/moc/models/record"
)

type ThemeController struct {
	beego.Controller
}

func (t *ThemeController) DoFinish() {
	recordThemeIDStr := t.Ctx.Input.Param(":id")
	recordThemeID, err := strconv.ParseInt(recordThemeIDStr, 10, 64)
	if err != nil {
		t.Abort(err.Error())
	}

	recordTheme, ok := record.GetTheme(recordThemeID)
	if !ok {
		t.Abort(fmt.Sprintf("No such record themd [%d]", recordThemeID))
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Abort(fmt.Sprintf("Record themd [%d] not in playing", recordThemeID))
	}

	recordTheme.Finish()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}

func (t *ThemeController) DoUnfinish() {
	recordThemeIDStr := t.Ctx.Input.Param(":id")
	recordThemeID, err := strconv.ParseInt(recordThemeIDStr, 10, 64)
	if err != nil {
		t.Abort(err.Error())
	}

	recordTheme, ok := record.GetTheme(recordThemeID)
	if !ok {
		t.Abort(fmt.Sprintf("No such record themd [%d]", recordThemeID))
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Abort(fmt.Sprintf("Record themd [%d] not in playing", recordThemeID))
	}

	recordTheme.Unfinish()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}

func (t *ThemeController) DoTip() {
	recordThemeIDStr := t.Ctx.Input.Param(":id")
	recordThemeID, err := strconv.ParseInt(recordThemeIDStr, 10, 64)
	if err != nil {
		t.Abort(err.Error())
	}

	recordTheme, ok := record.GetTheme(recordThemeID)
	if !ok {
		t.Abort(fmt.Sprintf("No such record themd [%d]", recordThemeID))
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Abort(fmt.Sprintf("Record themd [%d] not in playing", recordThemeID))
	}

	recordTheme.Tip()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}
