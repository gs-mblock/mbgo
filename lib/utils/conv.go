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

// StringToInt64Array : ("1,2,3",",")
func StringToInt64Array(idString string, delim string) []int64 {
	if idString == "" {
		return []int64{}
	}
	list := []int64{}
	s := strings.Split(idString, delim)
	for _, v := range s {
		intV, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
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

// StructToString :
func StructToString(i interface{}) string {
	out, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(out))
	return string(out)
}
