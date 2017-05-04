package theme

import (
	"log"

	"github.com/astaxie/beego/orm"
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
	_, err := orm.NewOrm().Update(t)
	if err != nil {
		panic(err)
	}
}

func GetThemesByStatus(status ThemeStatus) []Theme {
	themes := make([]Theme, 0)
	num, err := orm.NewOrm().QueryTable("theme").Filter("status", status).All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetThemesByStatus returned rows : %d ", num)
	return themes
}
