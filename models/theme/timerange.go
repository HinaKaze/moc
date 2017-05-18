package theme

import (
	"time"
)

type TimeRange struct {
	ID    int64
	Theme *Theme    `orm:"rel(fk)"`
	From  time.Time //开始时间，Hour和Minute有效
	To    time.Time //结束时间，Hour和Minute有效
}

func (t *TimeRange) TableName() string {
	return "theme_time_range"
}
