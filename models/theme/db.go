package theme

import (
	"log"

	"github.com/astaxie/beego/orm"
)

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

func InsertTheme(t *Theme) *Theme {
	var err error
	t.Id, err = orm.NewOrm().Insert(t)
	if err != nil {
		panic(err)
	}
	return t
}

func InsertTimeRange(timeRange *TimeRange) *TimeRange {
	var err error
	timeRange.Id, err = orm.NewOrm().Insert(timeRange)
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
	for index := range themes {
		_, err = orm.NewOrm().LoadRelated(&themes[index], "Tips")
		if err != nil {
			panic(err.Error())
		}

		_, err = orm.NewOrm().LoadRelated(&themes[index], "TimeRange")
		if err != nil {
			panic(err.Error())
		}
	}

	return themes
}
