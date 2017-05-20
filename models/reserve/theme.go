package reserve

import (
	"time"

	"github.com/hinakaze/moc/models/theme"
)

type ThemeStatus byte

const (
	ThemeStatusWaiting ThemeStatus = iota //预约中，等待开始
	ThemeStatusStarted                    //已经开始游戏
	ThemeStatusDeleted                    //手动删除
)

type Theme struct {
	Id          int64
	BeginTime   time.Time
	Theme       *theme.Theme `orm:"rel(fk)"`
	TeamName    string       //小队名称
	MemberCount int          //小队人数
	PhoneNumber string       //联系电话

	PayType  string  //支付类型
	PayPrice float64 //支付总金额

	Status ThemeStatus
}

func (t *Theme) TableName() string {
	return "reserve_theme"
}
