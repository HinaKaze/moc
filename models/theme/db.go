package theme

import (
	"log"

	"github.com/astaxie/beego/orm"
)

func GetTheme(id int64) (*Theme, bool) {
	t := new(Theme)
	t.ID = id
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

func InsertTheme(t *Theme) *Theme {
	var err error
	t.ID, err = orm.NewOrm().Insert(t)
	if err != nil {
		panic(err)
	}
	return t
}

func InsertTimeRange(timeRange *TimeRange) *TimeRange {
	var err error
	timeRange.ID, err = orm.NewOrm().Insert(timeRange)
	if err != nil {
		panic(err)
	}
	return timeRange
}

func UpdateTheme(t *Theme) {
	_, err := orm.NewOrm().Update(t)
	if err != nil {
		panic(err)
	}
}

func GetThemesByStatus(status ThemeStatus) []Theme {
	themes := make([]Theme, 0)
	num, err := orm.NewOrm().QueryTable(new(Theme).TableName()).Filter("status", status).All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetThemesByStatus returned rows : %d ", num)
	return themes
}
