package record

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/hinakaze/moc/models/reserve"
)

type Theme struct {
	Id      int64
	Reserve *reserve.Theme `orm:"rel(fk)"`

	RunningFlag  bool      //是否正在进行
	PayType      string    //支付类型
	PayPrice     float64   //支付总金额
	TipUsed      int       //使用提示数
	CompleteFlag bool      //是否破关
	TimeRecords  int       // seconds 破关记录
	CreateTime   time.Time //生成记录时间
	PlayDuration int       //游戏时长
}

func (t *Theme) TableName() string {
	return "record_theme"
}

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

func GetRunningRecords() []Theme {
	themes := make([]Theme, 0)
	num, err := orm.NewOrm().QueryTable("record_theme").Filter("running_flag", true).All(&themes)
	if err != nil {
		panic(err)
	}
	log.Printf("GetRunningRecords returned rows : %d ", num)
	return themes
}
