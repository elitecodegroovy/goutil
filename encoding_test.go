package goutil

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestEncoding(t *testing.T) {
	Convey("When generating base64 header", t, func() {
		result := GetBasicAuthHeader("grafana", "1234")

		So(result, ShouldEqual, "Basic Z3JhZmFuYToxMjM0")
	})

	Convey("When decoding basic auth header", t, func() {
		header := GetBasicAuthHeader("grafana", "1234")
		username, password, err := DecodeBasicAuthHeader(header)
		So(err, ShouldBeNil)

		So(username, ShouldEqual, "grafana")
		So(password, ShouldEqual, "1234")
	})

	Convey("When encoding password", t, func() {
		encodedPassword := EncodePassword("iamgod", "pepper")
		So(encodedPassword, ShouldEqual, "e59c568621e57756495a468f47c74e07c911b037084dd464bb2ed72410970dc849cabd71b48c394faf08a5405dae53741ce9")
	})
}

func TestStartFormatInt(t *testing.T) {
	strToInt := func(s string) int {
		//Atoi is shorthand for ParseInt(s, 10, 0).
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("error :", s, ", info :", err)
		}
		return num
	}
	a = "12"
	fmt.Println("output 14 is ", OuterFunc(strToInt, 2) == "14")
}
