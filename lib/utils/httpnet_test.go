package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestGetHTTPNetData :
func TestGetHTTPNetData(t *testing.T) {
	url := "https://api.xuebaclass.com/xuebaapi/v1/provinces"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type Province struct {
		ID       int    `json:"id"`
		Province string `json:"province"`
	}
	provinces := make([]Province, 0)
	err = json.Unmarshal([]byte(body), &provinces)
	if err != nil {
		println("error:", err)
	}
	fmt.Println(provinces)

}

// ---------------------
// 作者：一蓑烟雨1989
// 来源：CSDN
// 原文：https://blog.csdn.net/wangshubo1989/article/details/70245570
// 版权声明：本文为博主原创文章，转载请附上博文链接！
