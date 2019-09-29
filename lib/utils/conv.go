package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ArrayToString :
func ArrayToString(list []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(list), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

// StringToInt64Array : ("1,2,3",","),去重了
func StringToInt64Array(idString string, delim string) []int64 {
	if idString == "" {
		return []int64{}
	}
	list := []int64{}
	s := strings.Split(idString, delim)
	for _, v := range s {
		intV, err := strconv.ParseInt(v, 10, 64)
		if err == nil && !ArrayInt64Contains(list, intV) {
			list = append(list, intV)
		}
	}
	return list
}

// ArrayInt64Contains :array 是否存在某值
func ArrayInt64Contains(list []int64, item int64) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

// StructToString : MapToString
func StructToString(i interface{}) string {
	out, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(out))
	return string(out)
}

// BoolToInt :
func BoolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

// StringToMapList :
func StringToMapList(str string) ([]map[string]interface{}, error) {
	//str := `[{"status": "success", "type": "masterInfo", "result": "[{read: 2.0, write: 1.2}, {read_mb: 4.0, write: 3.2}]"}]`
	var m []map[string]interface{}
	err := json.Unmarshal([]byte(str), &m)
	if nil != err {
		//fmt.Println(err)
		return nil, err
	}
	//fmt.Println(m)
	return m, nil
}

// StringToMap :
func StringToMap(str string) (map[string]interface{}, error) {
	//str := `[{"status": "success", "type": "masterInfo", "result": "[{read: 2.0, write: 1.2}, {read_mb: 4.0, write: 3.2}]"}]`
	m := map[string]interface{}{}
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		//panic(err)
		return m, err
	}
	//fmt.Println(m)
	return m, nil
}

// StringToMapInterface :
func StringToMapInterface(str string) (interface{}, error) {
	s1 := string([]rune(str)[:1])
	//fmt.Println(s1)
	if s1 == "[" {
		return StringToMapList(str)
	}
	if s1 == "{" {
		return StringToMap(str)
	}
	return nil, nil
}

// MapToString :
// func MapToString(m map[string]string) string {
// 	b := new(bytes.Buffer)
// 	for key, value := range m {
// 		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
// 	}
// 	return b.String()
// }

// StringIsEmpty :
func StringIsEmpty(ss string) bool {
	//ss := strings.TrimSpace(s)
	//println("StringIsEmpty=", len(ss))
	//println("StringIsEmpty s=", (ss))
	if ss == "" || len(ss) == 0 {
		return true
	}
	return false
}

// Typeof :
func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// InterfaceToInt64 :
func InterfaceToInt64(v interface{}) int64 {
	println("InterfaceToInt64 V=", &v)
	println("InterfaceToInt64 ty=", Typeof(&v))
	value1, ok1 := v.(int)
	println("ok1=", ok1)
	println("value1=", value1)
	if ok1 {
		return int64(value1)
	}
	value2, ok2 := v.(string)
	if ok2 {
		VInt64, _ := strconv.ParseInt(value2, 10, 64)
		return VInt64
	}
	return 0
}

// InterfaceToInt642 :
// func InterfaceToInt642(v interface{}) int64 {
// 	type Bag struct {
// 		Key string
// 	}
// 	type Bag2 struct {
// 		Key int64
// 	}
// 	switch v := v.(type) { //v表示b1 接口转换成Bag对象的值
// 	case Bag:
// 		fmt.Println("b1.(type):", "Bag", v)
// 	case Bag2:
// 		fmt.Println("b1.(type):", "Bag2", v)
// 		return int64(v)
// 	default:
// 		return 0
// 	}

// }
