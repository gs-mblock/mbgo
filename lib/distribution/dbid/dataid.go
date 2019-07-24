package dbid

import "time"

/*
DataID ...
*分布式 id : 简化版
@businessID 业务ID 默认0 ,0-9;
@serviceID 服务器ID，默认0 ,0-9;
mysql.bigint= 9223372036854775807  19位数 -> 2262-04-12 07:47:16
			  1559787483440537000  19位数 -> 2019/06/06 10:34:58 【发现后3位是0】
可以用：2262 - 2019 = 243年 （实制2262年）
	//timestamp := time.Now().Unix()      // 1491888244 单位s,打印结果,10位数
	//timestamp := time.Now().UnixNano()  // 1559787483440537000 纳秒,19位数()
	//id := timestamp //+ int64(rand.Intn(100)*100) + (serviceID*10) + businessID  //结果：58608101
*/
// DataID 分布式 id : 简化版
func DataID(businessID int, serviceID int) int64 {
	id := (time.Now().UnixNano()-1546300800000000000)/1000*1000 + int64(businessID * 100) + int64(serviceID) //结果:3651562136279401  ,13+位数
	return id
}

// IDtoTime 把ID转化为时间
func IDtoTime(dbid int64) int64{
	timeid := (dbid + 1546300800000000000) / 1000000000
	return timeid
}