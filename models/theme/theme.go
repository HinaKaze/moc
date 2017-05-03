package theme

import (
	"log"

	"github.com/astaxie/beego/orm"
)

// func init(){
// 	fakeTheme1 := new(Theme)
// 	fakeTheme1.Title = "星际穿越"
// 	fakeTheme1.Desc = "一次去了就回不来的星际旅行"
// 	fakeTheme1.MinMember = 2
// 	fakeTheme1.MaxMember = 6
// 	fakeTheme1.PlayDuration = 3600
// 	fakeTheme1.Available = true

// 	InsertTheme(fakeTheme1)
// }

type Theme struct {
	Id           int64
	Title        string //主题名称
	Desc         string //主题说明
	MinMember    int    //最小参与人数
	MaxMember    int    //最大参与人数
	PlayDuration int    //规定游玩时长 seconds
	Available    bool   //当前是否有效
}

func (t *Theme) TableName() string {
	return "theme"
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

func GetAvailableThemes() []Theme {
	themes := make([]Theme, 0)
	num, err := orm.NewOrm().QueryTable("theme").Filter("available", true).All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetThemes returned rows : %d ", num)
	return themes
}
