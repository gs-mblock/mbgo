package utils

import (
	"fmt"
	"strconv"
	"testing"
	"time"
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

	// int64 to int
	int0 := (int32)(time.Now().Unix())
	println("int0", int0)
}

func Test_ArrayToString(t *testing.T) {
	list := []int64{1, 2, 3, 4}
	//list := []int64{1}
	rs := ArrayToString(list, ",")
	fmt.Println(rs)
	// 1,2,3,4
}

func Test_StringToInt64Array(t *testing.T) {
	list := StringToInt64Array("1,2,3", ",")
	for i, v := range list {
		println(i, "--", v)
	}
}
