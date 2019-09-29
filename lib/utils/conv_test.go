package utils

import (
	"encoding/json"
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
	list := StringToInt64Array("1,1,2,3,3", ",")
	for i, v := range list {
		println(i, "--", v)
	}
}

func Test2_interface(t *testing.T) {
	type Bag struct {
		Key string
	}

	type Bag2 struct {
		Key int64
	}

	var b1 interface{}
	var b2 interface{}

	b1 = Bag{Key: "1"}
	b2 = Bag2{Key: 0}
	//获取interface{}中存放的数据类型
	//方法一：
	{ //判断是否是Bag类型  若不是则置0
		b, ok := b1.(Bag)
		fmt.Println("Bag类型   ：", ok, "数据", b)
	}
	{ //判断是否是Bag2类型  若不是则置0
		b, ok := b2.(Bag2)
		fmt.Println("Bag2类型：", ok, "数据", b)
	}

}

func TestIn_interface(t *testing.T) {
	var x interface{}
	x = "2392"
	vv := fmt.Sprintf("%v", x)

	println("vv=", vv)
	//println(InterfaceToInt64(x))

	value, ok := x.(int)
	value2, ok2 := x.(string)
	println("ok:", ok)
	println("value:", value)

	println("ok2:", ok2)
	println("value2=", value2)
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

func TestConvent_StringToMap(t *testing.T) {
	str := `[{"status": "success", "type": "masterInfo", "result": "[{read: 2.0, write: 1.2}, {read_mb: 4.0, write: 3.2}]"}]`
	var m []map[string]interface{}
	err := json.Unmarshal([]byte(str), &m)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
}

func TestConvent_StringToMapList(t *testing.T) {
	str := `[{"status": "success", "type": "masterInfo", "result": "[{read: 2.0, write: 1.2}, {read_mb: 4.0, write: 3.2}]"}]`
	db, err := StringToMapList(str)
	fmt.Println(err)
	fmt.Printf("%+v\n", db)
}

func TestConvent_MapToString(t *testing.T) {
	m := map[string]string{
		"LOG_LEVEL": "DEBUG",
		"API_KEY":   "12345678-1234-1234-1234-1234-123456789abc",
	}
	println(StructToString(m))
}

func TestConvent_StringToMap2(t *testing.T) {
	str1 := `[{"status": "success", "type": "masterInfo", "result": "[{read: 2.0, write: 1.2}, {read_mb: 4.0, write: 3.2}]"}]`
	str2 := `{"status": "success", "type": "masterInfo"}`

	// sec := map[string]interface{}{}
	// if err := json.Unmarshal([]byte(str1), &sec); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", sec)

	//s := "a我cd"
	s1 := string([]rune(str1)[:1])
	fmt.Println(s1)
	if s1 == "[" {
		vv, _ := StringToMapList(str1)
		fmt.Printf("%+v\n", vv)
	}

	s2 := string([]rune(str2)[:1])
	fmt.Println(s2)
	if s2 == "{" {
		vv, _ := StringToMap(str2)
		fmt.Printf("%+v\n", vv)
	}
}
