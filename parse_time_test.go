package goutil

import (
	"fmt"
	"testing"
	"time"
)

func TestParseIn(t *testing.T) {
	denverLoc, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		t.Fatal(err.Error())
	}
	var time1 time.Time
	if time1, err = ParseIn("3/1/2014", denverLoc); err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("time1: %+v \n", time1)
	fmt.Printf("time1: %+v \n", GetISOStrTime(time1))

	//     denverLoc, _ := time.LoadLocation("Asia/Chongqing)
	//     time.Local = denverLoc
	//
	//     t, err := dateparse.ParseLocal("3/1/2014")
	//
	// Equivalent to:
	//
	//     t, err := dateparse.ParseIn("3/1/2014", denverLoc)
}
