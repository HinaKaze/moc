package record

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/hinakaze/moc/models/record"
)

type ThemeController struct {
	beego.Controller
}

func (t *ThemeController) Finish() {
	recordThemeIdStr := t.Ctx.Input.Param(":id")
	recordThemeId, err := strconv.ParseInt(recordThemeIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	recordTheme, ok := record.GetTheme(recordThemeId)
	if !ok {
		t.Redirect("/err", 500)
		return
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Redirect("/err", 500)
		return
	}

	recordTheme.Finish()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}

func (t *ThemeController) Unfinish() {
	recordThemeIdStr := t.Ctx.Input.Param(":id")
	recordThemeId, err := strconv.ParseInt(recordThemeIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	recordTheme, ok := record.GetTheme(recordThemeId)
	if !ok {
		t.Redirect("/err", 500)
		return
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Redirect("/err", 500)
		return
	}

	recordTheme.Unfinish()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}

func (t *ThemeController) Tip() {
	recordThemeIdStr := t.Ctx.Input.Param(":id")
	recordThemeId, err := strconv.ParseInt(recordThemeIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	recordTheme, ok := record.GetTheme(recordThemeId)
	if !ok {
		t.Redirect("/err", 500)
		return
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Redirect("/err", 500)
		return
	}

	recordTheme.Tip()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}
