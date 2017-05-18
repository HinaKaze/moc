package main

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hinakaze/iniparser"
	_ "github.com/hinakaze/moc/routers"
	_ "github.com/lib/pq"

	"github.com/hinakaze/moc/models/record"
	"github.com/hinakaze/moc/models/reserve"
	"github.com/hinakaze/moc/models/theme"
)

func main() {
	/*
		db init
	*/
	iniparser.DefaultParse("./conf/user.ini")
	section, ok := iniparser.GetSection("DB")
	if !ok {
		panic("ini parse error")
	}
	driverName, ok := section.GetValue("driverName")
	if !ok {
		panic("[driverName] not found")
	}
	dataSource, ok := section.GetValue("dataSource")
	if !ok {
		panic("[dataSource] not found")
	}

	orm.Debug = true
	orm.RegisterDataBase("default", driverName, dataSource)
	orm.DefaultTimeLoc = time.Local

	orm.RegisterModel(new(theme.Theme), new(reserve.Theme), new(record.Theme), new(theme.Tip), new(theme.TimeRange))

	//beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.Run()
}
