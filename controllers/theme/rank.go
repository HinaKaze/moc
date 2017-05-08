package theme

import (
	"strconv"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/hinakaze/moc/models/record"
	"github.com/hinakaze/moc/models/theme"
)

type RankController struct {
	beego.Controller
}

func (t *RankController) GetRank() {
	themeIDStr := t.Ctx.Input.Param(":id")
	themeID, err := strconv.ParseInt(themeIDStr, 10, 64)
	if err != nil {
		t.Abort(err.Error())
	}

	theme, ok := theme.GetTheme(themeID)
	if !ok {
		t.Abort(fmt.Sprintf("No such theme [%d]", themeID))
	}

	rank := record.GetThemesRank(themeID, 10)

	t.Data["Rank"] = rank
	t.Data["Theme"] = theme
	t.TplName = "theme/rank.html"
}
