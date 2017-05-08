package record

import (
	"log"

	"github.com/astaxie/beego/orm"
)

func GetTheme(id int64) (*Theme, bool) {
	theme := new(Theme)
	theme.Id = id
	err := orm.NewOrm().Read(theme)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, false
		} else {
			panic(err)
		}
	}
	return theme, true
}

func InsertTheme(t *Theme) *Theme {
	var err error
	t.Id, err = orm.NewOrm().Insert(t)
	if err != nil {
		panic(err)
	}
	return t
}

func UpdateTheme(t *Theme) {
	_, err := orm.NewOrm().Update(t)
	if err != nil {
		panic(err)
	}
}

func GetThemesByStatus(status ThemeStatus) []Theme {
	themes := make([]Theme, 0)
	num, err := orm.NewOrm().QueryTable(new(Theme).TableName()).Filter("status", status).RelatedSel().All(&themes)
	if err != nil {
		panic(err)
	}

	log.Printf("GetRecordsByStatus returned rows : %d ", num)
	return themes
}

/*
rule :
	1 . finished
	2 . used time
	3 . begin time
*/

func GetThemesRank(themeID int64, limit int) []Theme {
	themes := make([]Theme, 0)
	num, err := orm.NewOrm().QueryTable(new(Theme).TableName()).Filter("status", ThemeStatusFinished).Filter("reserve__theme__id", themeID).OrderBy("time_used", "begin_time").Limit(limit).RelatedSel().All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetThemesRank themeID [%d] limit [%d] returned rows : %d", themeID, limit, num)
	return themes
}
