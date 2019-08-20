package utils

import (
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