package utils

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestConvString2Other(t *testing.T) {
	VString := "123"
	//# string到int
	VInt, err := strconv.Atoi(VString)
	if err != nil {
		println("string到int:", err)
	}
	//#string到int64
	VInt64, err := strconv.ParseInt(VString, 10, 64)
	//#int到string
	VString = strconv.Itoa(VInt)
	//#int64到string
	VString = strconv.FormatInt(VInt64, 10)
	println("int64到string:", VString)

	// int64 to int
	int0 := (int32)(time.Now().Unix())
	println("int64 to int:", int0)

	// int to int64
	VI := 23
	var i64 int64
	i64 = int64(VI)
	println("int to int64:", i64)

	// bool to string
	sb := strconv.FormatBool(true)
	println("bool to string:", sb)

	// string to byte
	str := "test"
	dataByte := []byte(str)
	fmt.Println("byte:", dataByte)
	// byte to string
	strByte := string(dataByte[:])
	fmt.Println("st:", strByte)
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

func TestConversion_ArrayInt64Contains(t *testing.T) {
	data1 := ArrayInt64Contains([]int64{1, 2, 3}, 3)
	println("data1:", data1)
	data2 := ArrayInt64Contains([]int64{}, 1)
	println("data2:", data2)
}

func TestStructToString(t *testing.T) {
	type Server struct {
		Name    string `json:"name,omitempty"`
		ID      int
		Enabled bool
		users   []string // not exported
		//http.Server          // embedded
	}
	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}
	s := StructToString(server)
	println(s)
}
