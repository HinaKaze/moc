package theme

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/hinakaze/moc/models/record"
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

	rank := record.GetThemesRank(themeID, 10)

	t.Data["Rank"] = rank
	t.TplName = "theme/rank.html"
}
