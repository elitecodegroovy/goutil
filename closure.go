package goutil

import (
	"fmt"
	"strconv"
	"testing"
)

var a string = "1"

func OuterFunc(strToInt func(s string) int, b int) string {
	c := strToInt(a) + b
	//int is converted to string type
	return strconv.Itoa(c)
}
