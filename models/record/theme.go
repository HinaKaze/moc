package record

import (
	"time"

	"github.com/hinakaze/moc/models/reserve"
)

type ThemeStatus byte

const (
	ThemeStatusPlaying    ThemeStatus = iota //正在游戏
	ThemeStatusUnfinished                    //游戏结束，且未成功破关
	ThemeStatusFinished                      //游戏结束，且成功破关
)

type Theme struct {
	Id      int64
	Reserve *reserve.Theme `orm:"rel(fk)"`

	PayType   string    //支付类型
	PayPrice  float64   //支付总金额
	TipUsed   int       //使用提示数
	TimeUsed  int       //游戏时长，seconds
	BeginTime time.Time //游戏开始时间

	Status ThemeStatus
}

func (t *Theme) TableName() string {
	return "record_theme"
}

//Tip 使用一次提示
func (t *Theme) Tip() {
	t.TipUsed++
}

//Finish 破关结算
func (t *Theme) Finish() {
	timeUsed := time.Now().Sub(t.BeginTime) / time.Second
	t.TimeUsed = int(timeUsed)
	t.Status = ThemeStatusFinished
}

//Unfinish 未破关结算
func (t *Theme) Unfinish() {
	timeUsed := time.Now().Sub(t.BeginTime) / time.Second
	t.TimeUsed = int(timeUsed)
	t.Status = ThemeStatusUnfinished
}

func CreateTheme(reserveTheme *reserve.Theme, payType string, payPrice float64) *Theme {
	t := new(Theme)
	t.Reserve = reserveTheme
	t.PayType = payType
	t.PayPrice = payPrice
	t.TipUsed = 0
	t.TimeUsed = 0
	t.BeginTime = time.Now()
	t.Status = ThemeStatusPlaying
	return t
}
