package record

import (
	"log"
	"strconv"

	"runtime/debug"

	"github.com/astaxie/beego"
	"github.com/hinakaze/moc/models/record"
)

type ThemeController struct {
	beego.Controller
}

func (t *ThemeController) DoFinish() {
	log.Printf("%s", debug.Stack())
	recordThemeIdStr := t.Ctx.Input.Param(":id")
	recordThemeId, err := strconv.ParseInt(recordThemeIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	recordTheme, ok := record.GetTheme(recordThemeId)
	if !ok {
		t.Abort("500")
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Abort("500")
	}

	recordTheme.Finish()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}

func (t *ThemeController) DoUnfinish() {
	recordThemeIdStr := t.Ctx.Input.Param(":id")
	recordThemeId, err := strconv.ParseInt(recordThemeIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	recordTheme, ok := record.GetTheme(recordThemeId)
	if !ok {
		t.Abort("500")
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Abort("500")
	}

	recordTheme.Unfinish()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}

func (t *ThemeController) DoTip() {
	recordThemeIdStr := t.Ctx.Input.Param(":id")
	recordThemeId, err := strconv.ParseInt(recordThemeIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	recordTheme, ok := record.GetTheme(recordThemeId)
	if !ok {
		t.Abort("500")
	}
	if recordTheme.Status != record.ThemeStatusPlaying {
		t.Abort("500")
	}

	recordTheme.Tip()
	record.UpdateTheme(recordTheme)
	t.Redirect("/", 302)
}
