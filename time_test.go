package goutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func TestGetISOStrDateTime(t *testing.T) {
	Convey("GetISOStrDateTime", t, func() {
		layout := "2006-01-02"
		str := "2019-11-12"
		tt1, err := time.Parse(layout, str)
		if err != nil {
			t.Error("time parse error:" + err.Error())
			os.Exit(1)
		}
		So(GetISOStrDate(tt1), ShouldEqual, str)

	})
}

func TestGetCurrentISOStrDateTime(t *testing.T) {
	Convey("GetCurrentDateISOStrTime", t, func() {
		t1 := time.Now()
		currentStrTime := t1.Format("2006-01-02")
		So(GetCurrentDateISOStrDate(), ShouldEqual, currentStrTime)
	})
}

func TestGetISOStrTime(t *testing.T) {
	Convey("GetISOStrTime", t, func() {
		layout := "2006-01-02 15:04:05"
		str := "2019-11-12 00:00:00"
		tt1, _ := time.Parse(layout, str)

		So(GetISOStrTime(tt1), ShouldEqual, str)

		str = "2019-11-12 10:10:20"
		tt1, _ = time.Parse(layout, str)
		So(GetISOStrTime(tt1), ShouldEqual, str)

		t2 := time.Now()
		t2str := GetISOStrTime(t2)
		So(GetCurrentTimeISOStrTime(), ShouldEqual, t2str)
	})
}

func TestGetISOStrTimeNumber(t *testing.T) {
	Convey("GetISOStrTimeNumber", t, func() {
		layout := "20060102150405"
		str := "20191112000000"
		tt1, _ := time.Parse(layout, str)

		So(GetISOStrTimeNumber(tt1), ShouldEqual, str)

		str = "20191112101020"
		tt1, _ = time.Parse(layout, str)
		So(GetISOStrTimeNumber(tt1), ShouldEqual, str)

		t2 := time.Now()
		t2str := GetISOStrTimeNumber(t2)
		So(GetCurrentTimeNumberISOStrTime(), ShouldEqual, t2str)
	})
}
