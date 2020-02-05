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

func TestEncodeMd5(t *testing.T) {
	Convey("md5 encoding", t, func() {
		encodedPassword := EncodeMd5("123456")
		So(encodedPassword, ShouldEqual, "e10adc3949ba59abbe56e057f20f883e")
	})
}

func TestEncodeSha1(t *testing.T) {
	Convey("md5 encoding", t, func() {
		encodedPassword := EncodeSha1("sha1 this string")
		So(encodedPassword, ShouldEqual, "cf23df2207d99a74fbe169e3eba035e633b65d94")
	})
}

func TestEncodeSha256(t *testing.T) {
	Convey("md5 encoding", t, func() {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{[]byte(""), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
			{[]byte("abc"), "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
			{[]byte("hello"), "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		} {
			t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
				if EncodeSha256(string(tt.in)) != tt.out {
					t.Errorf("want %v; got %v", tt.out, EncodeSha256(string(tt.in)))
				}
			})
		}
	})
}

func TestEncodeSha512(t *testing.T) {
	Convey("md5 encoding", t, func() {
		encodedPassword := EncodeSha512("sha512 this string")
		So(encodedPassword, ShouldEqual, "3eb8ad2add74c22fc006851058a39a74b73dc5f6eadb0fefb829c5fa4572faffabfe3df7cff7baa62fab280b153c0bbd99d317737305d59ac89c8114dc8139b6")
	})
}
