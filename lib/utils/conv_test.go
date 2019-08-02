package utils

import (
	"strconv"
	"testing"
)

func TestConvString2Other(t *testing.T) {
	string := "123"
	//# string到int
	int, err := strconv.Atoi(string)
	if err != nil {
		println(err)
	}
	//#string到int64
	int64, err := strconv.ParseInt(string, 10, 64)
	//#int到string
	string = strconv.Itoa(int)
	//#int64到string
	string = strconv.FormatInt(int64, 10)
	println(string)
}
