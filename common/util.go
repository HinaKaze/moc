package common

import (
	"time"

	"github.com/astaxie/beego"
)

var TimeLocation *time.Location

func init() {
	var err error
	TimeLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

//FormatTime 转换time.Time to string
func FormatTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

//ParseTime 解析string to time.Time
func ParseTime(str string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", str, TimeLocation)
}

func HandleSuccess(r *beego.Controller) {
	if r == nil {
		panic("HandleSuccess occur panic , controller is nil")
	}
	var json struct {
		result int
	}
	json.result = 1
	r.Data["json"] = &json
	r.ServeJSON()
}
