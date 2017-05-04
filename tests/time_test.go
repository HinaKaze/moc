package test

import (
	"log"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	timeStr := "2017-05-24 10:24:58"
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05 +0800 CST", timeStr, time.Local)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(theTime.String())
}
