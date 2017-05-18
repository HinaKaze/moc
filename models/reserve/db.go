package reserve

import (
	"log"

	"time"

	"github.com/astaxie/beego/orm"
	"github.com/hinakaze/moc/common"
)

func InsertTheme(t *Theme) *Theme {
	var err error
	t.Id, err = orm.NewOrm().Insert(t)
	if err != nil {
		panic(err)
	}
	return t
}

func UpdateTheme(t *Theme) {
	num, err := orm.NewOrm().Update(t)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("UpdateTheme updates rows : %d ", num)
}

func GetTheme(id int64) (*Theme, bool) {
	t := new(Theme)
	t.Id = id
	err := orm.NewOrm().Read(t)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, false
		} else {
			panic(err)
		}
	}
	return t, true
}

func GetThemesByStatus(status ThemeStatus) []Theme {
	themes := make([]Theme, 0)
	num, err := orm.NewOrm().QueryTable(new(Theme).TableName()).Filter("status", status).OrderBy("begin_time").RelatedSel().All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetThemesByStatus returned rows : %d ", num)
	return themes
}

func GetOneDayThemes(day time.Time) []Theme {
	from := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, common.TimeLocation)
	to := from.AddDate(0, 0, 1)
	return GetHistoryThemes(from, to)
}

func GetHistoryThemes(from time.Time, to time.Time) []Theme {
	themes := make([]Theme, 0)

	num, err := orm.NewOrm().QueryTable(new(Theme).TableName()).Filter("begin_time__gte", from).Filter("begin_time__lt", to).RelatedSel().All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetHistoryThemes returned rows : %d", num)
	return themes
}
