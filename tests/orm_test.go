package test

import (
	"fmt"
	"testing"

	"time"

	"github.com/astaxie/beego/orm"
	"github.com/hinakaze/iniparser"
	"github.com/hinakaze/moc/models/record"
	"github.com/hinakaze/moc/models/reserve"
	"github.com/hinakaze/moc/models/theme"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	fmt.Println("Start beego orm test")
	iniparser.DefaultParse("../conf/user.ini")
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

	m.Run()
}

func TestCreateModels(t *testing.T) {
	orm.RegisterModel(new(theme.Theme), new(reserve.Theme), new(record.Theme))
	orm.RunSyncdb("default", false, true)
}

func TestInsertThemeData(t *testing.T) {
	fakeTheme1 := new(theme.Theme)
	fakeTheme1.Title = "星际穿越"
	fakeTheme1.Desc = "一次去了就回不来的星际旅行"
	fakeTheme1.MinMember = 2
	fakeTheme1.MaxMember = 6
	fakeTheme1.PlayDuration = 3600
	fakeTheme1.Status = theme.ThemeStatusOpening

	fakeTheme2 := new(theme.Theme)
	fakeTheme2.Title = "电锯惊魂2"
	fakeTheme2.Desc = "也就少条腿什么的"
	fakeTheme2.MinMember = 4
	fakeTheme2.MaxMember = 10
	fakeTheme2.PlayDuration = 3600
	fakeTheme2.Status = theme.ThemeStatusOpening

	theme.InsertTheme(fakeTheme1)
	theme.InsertTheme(fakeTheme2)
}

func TestInsertReserveData(t *testing.T) {
	now := time.Now()
	themes := theme.GetThemesByStatus(theme.ThemeStatusOpening)
	for _, t := range themes {
		for i := 10; i <= 20; i++ {
			fakeReserve := new(reserve.Theme)
			fakeReserve.Theme = &t
			fakeReserve.TeamName = "紫辰战队"
			fakeReserve.MemberCount = i - 8
			fakeReserve.PhoneNumber = "123456789"
			fakeReserve.BeginTime = time.Date(now.Year(), now.Month(), now.Day(), i, 10, 0, 0, time.Local)
			fakeReserve.Status = reserve.ThemeStatusWaiting
			reserve.InsertTheme(fakeReserve)
		}
	}
}
