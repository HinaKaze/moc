package reserve

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
	// cond := orm.NewCondition()
	// cond.And("BeginTime_gt", time.Now())
	//num, err := orm.NewOrm().QueryTable("reserve_theme").Filter("status", Status).SetCond(cond).All(&themes)

	num, err := orm.NewOrm().QueryTable(new(Theme).TableName()).Filter("status", status).RelatedSel().All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetThemesByStatus returned rows : %d ", num)
	return themes
}
