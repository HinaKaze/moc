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

	orm.RegisterModel(new(theme.Theme), new(reserve.Theme), new(record.Theme), new(theme.Tip), new(theme.TimeRange))

	m.Run()
}

func TestCreateModels(t *testing.T) {
	orm.RunSyncdb("default", false, true)
}

func TestInsertThemeData(t *testing.T) {
	fakeTheme1 := new(theme.Theme)
	fakeTheme1.Title = "星际救援"
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

	fakeTheme3 := new(theme.Theme)
	fakeTheme3.Title = "咒怨2"
	fakeTheme3.Desc = "也就少条腿什么的"
	fakeTheme3.MinMember = 4
	fakeTheme3.MaxMember = 10
	fakeTheme3.PlayDuration = 3600
	fakeTheme3.Status = theme.ThemeStatusOpening

	theme.InsertTheme(fakeTheme1)
	theme.InsertTheme(fakeTheme2)
	theme.InsertTheme(fakeTheme3)
}

func TestInsertThemeTimeRange(t *testing.T) {
	now := time.Now()
	themes := theme.GetThemesByStatus(theme.ThemeStatusOpening)
	for _, t := range themes {
		for i := 10; i <= 20; i++ {
			fakeTimeRange := new(theme.TimeRange)
			fakeTimeRange.Theme = &t
			fakeTimeRange.From = time.Date(now.Year(), now.Month(), now.Day(), i, 10, 0, 0, time.Local)
			fakeTimeRange.To = time.Date(now.Year(), now.Month(), now.Day(), i+1, 9, 0, 0, time.Local)
			theme.InsertTimeRange(fakeTimeRange)
		}
	}
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

func TestInsertTipData(t *testing.T) {
	themes := theme.GetThemesByStatus(theme.ThemeStatusOpening)
	for _, t := range themes {
		for i := 1; i <= 3; i++ {
			fakeTip := new(theme.Tip)
			fakeTip.Theme = &t
			fakeTip.Stage = fmt.Sprintf("Stage:%d", i)
			fakeTip.Desc = "自己想！你有脑子吗"
			theme.InsertTip(fakeTip)
		}
	}
}
