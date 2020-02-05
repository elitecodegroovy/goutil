package goutil

import (
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringsUtil(t *testing.T) {
	Convey("Falling back until none empty string", t, func() {
		So(StringsFallback2("1", "2"), ShouldEqual, "1")
		So(StringsFallback2("", "2"), ShouldEqual, "2")
		So(StringsFallback3("", "", "3"), ShouldEqual, "3")
	})
}

func TestSplitString(t *testing.T) {
	Convey("Splits strings correctly", t, func() {
		So(SplitString(""), ShouldResemble, []string{})
		So(SplitString("test"), ShouldResemble, []string{"test"})
		So(SplitString("test1 test2 test3"), ShouldResemble, []string{"test1", "test2", "test3"})
		So(SplitString("test1,test2,test3"), ShouldResemble, []string{"test1", "test2", "test3"})
		So(SplitString("test1, test2, test3"), ShouldResemble, []string{"test1", "test2", "test3"})
		So(SplitString("test1 , test2 test3"), ShouldResemble, []string{"test1", "test2", "test3"})
	})
}

func TestDateAge(t *testing.T) {
	Convey("GetAgeString", t, func() {
		So(GetAgeString(time.Time{}), ShouldEqual, "?")
		So(GetAgeString(time.Now().Add(-time.Second*2)), ShouldEqual, "< 1m")
		So(GetAgeString(time.Now().Add(-time.Minute*2)), ShouldEqual, "2m")
		So(GetAgeString(time.Now().Add(-time.Hour*2)), ShouldEqual, "2h")
		So(GetAgeString(time.Now().Add(-time.Hour*24*3)), ShouldEqual, "3d")
		So(GetAgeString(time.Now().Add(-time.Hour*24*67)), ShouldEqual, "2M")
		So(GetAgeString(time.Now().Add(-time.Hour*24*409)), ShouldEqual, "1y")
	})
}

func TestGetISOStrDateTime(t *testing.T) {
	Convey("GetISOStrDateTime", t, func() {
		layout := "2006-01-02"
		str := "2019-11-12"
		tt1, err := time.Parse(layout, str)
		if err != nil {
			t.Error("time parse error:" + err.Error())
			os.Exit(1)
		}
		So(GetISOStrDateTime(tt1), ShouldEqual, str)

	})
}

func TestGetCurrentISOStrDateTime(t *testing.T) {
	Convey("GetCurrentDateISOStrTime", t, func() {
		t1 := time.Now()
		currentStrTime := t1.Format("2006-01-02")
		So(GetCurrentDateISOStrTime(), ShouldEqual, currentStrTime)
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
