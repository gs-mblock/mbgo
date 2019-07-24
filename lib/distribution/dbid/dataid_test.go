package dbid

/**
cd framework/distribution
go test
*/
import (
	"testing"
	"time"
)

func Test_DataID(t *testing.T) {
	//test 
	mid  := DataID(4,1)
	println("[Test] get value", mid)
	if mid  > 0 {
		t.Log("Test_DataID 通过")
	}
}

func Test_IDtoTime(t *testing.T) {
	//test 
	mid  := DataID(4,1)
	println("[Test] get value", mid)
	if mid  > 0 {
		t.Log("Test_DataID 通过")
	}
	timeid := IDtoTime(mid)
	formatTimeStr1:=time.Unix(timeid,0).Format("2006-01-02 15:04:05")
	println("time S1:",formatTimeStr1)
}

func Test_Time2(t *testing.T) {
	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00") // 0时区
	t1 := tm2.Unix() // 1546300800
	t2 := tm2.UnixNano() // 1546300800000000000
	println("time U1:",t1)
	println("time U2:",t2)
	formatTimeStr1:=time.Unix(t1,0).Format("2006-01-02 15:04:05")
	println("time S1:",formatTimeStr1) // 2019-01-01 08:00:00
}

func Test_Time(t *testing.T) {
	t1 := time.Now().Unix()
	t2 := time.Now().UnixNano()
	tu2 := time.Now().UTC().Unix() // 时间戳 不分时区
	println("time U0:",tu2)
	println("time U1:",t1)

	println("time N0:",t2)
	formatTimeStr1:=time.Unix(t1,0).Format("2006-01-02 15:04:05")
	formatTimeStr2:=time.Unix(t1,0).UTC().Format("2006-01-02 15:04:05") // 0时区
	println("time S1:",formatTimeStr1)
	println("time S2:",formatTimeStr2)

	tz1 :=  time.Now().UTC().Format("2006-01-02 15:04:05") // 0时区
	tz2 :=  time.Now().Format("2006-01-02 15:04:05") 
	println("time Z1:",tz1)
	println("time Z1:",tz2)
}
