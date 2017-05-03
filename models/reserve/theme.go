package reserve

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/hinakaze/moc/models/theme"
)

type Theme struct {
	Id          int64
	Theme       *theme.Theme `orm:"rel(fk)"`
	TeamName    string       //小队名称
	BeginTime   time.Time    //游玩开始时间
	MemberCount int          //小队人数
	PhoneNumber string       //联系电话
}

func (t *Theme) TableName() string {
	return "reserve_theme"
}

func InsertTheme(t *Theme) *Theme {
	var err error
	t.Id, err = orm.NewOrm().Insert(t)
	if err != nil {
		panic(err)
	}
	return t
}

func GetTodayReserves() []Theme {
	themes := make([]Theme, 0)
	cond := orm.NewCondition()
	cond.And("BeginTime_gt", time.Now())
	num, err := orm.NewOrm().QueryTable("reserve_theme").Filter("running_flag", false).SetCond(cond).All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetTodayReserves returned rows : %d ", num)
	return themes
}
