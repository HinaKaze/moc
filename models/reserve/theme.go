package reserve

import (
	"time"

	"github.com/hinakaze/moc/models/theme"
)

type ThemeStatus byte

const (
	ThemeStatusWaiting   ThemeStatus = iota //预约中，等待开始
	ThemeStatusConverted                    //正常开始，并转换成为一个Record
	ThemeStatusExpired                      //过期但没来玩
	ThemeStatusDeleted                      //手动删除
)

type Theme struct {
	Id          int64
	Theme       *theme.Theme `orm:"rel(fk)"`
	TeamName    string       //小队名称
	BeginTime   time.Time    //游玩开始时间
	MemberCount int          //小队人数
	PhoneNumber string       //联系电话

	Status ThemeStatus
}

func (t *Theme) TableName() string {
	return "reserve_theme"
}
