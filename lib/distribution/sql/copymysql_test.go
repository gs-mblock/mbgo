package sql

import (
	"testing"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"gitlab.com/makeblock-common/mbgo/lib/utils"
	"gitlab.com/makeblock-common/mbgo/lib/distribution/dbid"

	"strconv"
	"bytes"
)
type User struct {
    UserID int64  `xorm:"user_id"`
    Name string  `xorm:"name"`
    CreateTime int `xorm:"create_time"`
	ModifyTime int `xorm:"modify_time"`
}

func TestInit(t *testing.T){
	url1 := "root:1234qwer@(localhost:3306)/db_test1?charset=utf8"
	url2 := "root:1234qwer@(localhost:3306)/db_test2?charset=utf8"
    engine1, err := xorm.NewEngine("mysql", url1)
    if err != nil{
		fmt.Println(err)
		t.Fatal(err)
	}
	engine2, err := xorm.NewEngine("mysql", url2)
	//println(engine)
	creatTable(engine1)
	creatTable(engine2)
	engine2.Close()
	engine1.Close()
}

// 生产测试数据
func TestInsertData(t *testing.T){
	url1 := "root:1234qwer@(localhost:3306)/db_test1?charset=utf8"
	url2 := "root:1234qwer@(localhost:3306)/db_test2?charset=utf8"
    engine1, err := xorm.NewEngine("mysql", url1)
    if err != nil{
		fmt.Println(err)
		t.Fatal(err)
	}
	engine2, err := xorm.NewEngine("mysql", url2)
	for a := 0; a < 10; a++ {
		insert(engine1,1)
		insert(engine2,2)
	}
	engine2.Close()
	engine1.Close()
}

// 测试同步数据
func TestMysqlSynchronousData(t *testing.T){
	url1 := "root:1234qwer@(localhost:3306)/db_test1?charset=utf8"
	url2 := "root:1234qwer@(localhost:3306)/db_test2?charset=utf8"
    engine1, err := xorm.NewEngine("mysql", url1)
    if err != nil{
		fmt.Println(err)
		t.Fatal(err)
	}
	engine2, err := xorm.NewEngine("mysql", url2)
	println(engine1)
	println(engine2)
	ids1 := getAllID(engine1)
	ids2 := getAllID(engine2)
	add1 := checkIDGetID1News(ids1,ids2)
	println(add1)
	add2 := checkIDGetID1News(ids2,ids1)
	println(add2)
	// db1  添加 db2 数据； 
	addData2ToData1(engine1,engine2,add2)
	// db2 添加 db1 数据； 
	addData2ToData1(engine2,engine1,add1)
}



func creatTable(engine *xorm.Engine){
	engine.DropTables(new(User))
    err := engine.CreateTables(new(User))
    if err != nil{
        fmt.Println("error in create table user, ", err)
	}
}

func insert(engine *xorm.Engine,serverID int){
	var user User
	user.UserID = dbid.DataID(4,serverID)
	user.Name = "sv-"+strconv.Itoa(serverID)+ "_"+ string( utils.RandomCreateBytes(5))
	user.CreateTime = int(time.Now().Unix()) 
	user.ModifyTime = int(time.Now().Unix() )
	_, err := engine.Insert(&user)
	if err != nil{
		fmt.Println(err)
	}
	println("[Test]", user.UserID)
}

func getAllID(engine *xorm.Engine) (ids []int64) {
	// 同步 2分钟前的数据
	t0 := time.Now().Unix() - (60*60*24)
	t1 := time.Now().Unix() - (60*2)
	sql := fmt.Sprintf("select user_id from user where modify_time >= %d and modify_time <= %d ", t0, t1)
	println("[TEST]sql=",sql)
	results, err := engine.QueryString(sql)
	if err != nil{
		println(err)
	}
	
	for i, v := range results{
		//println(k)
		for m, uid := range v{
			println(i, "[TEST] get:",m,uid )
			//println(i)
			ii, _ := strconv.ParseInt(uid, 10, 64)
			ids = append(ids, ii)
		}
	}
	return ids
}

// 找出 id1 中，新加的数据
func checkIDGetID1News(id1 []int64,id2 []int64) (a1 []int64){
	for _, v1 := range id1 {
		add :=1
		for _, v2 := range id2 {
			if v1 == v2 {
				add =0
				break
			}
		}
		if add == 1 {
			a1 = append(a1, v1)
			println("[TEST] add ",v1)
		} 
	}
	return a1
}

// 从 db2,拉数据到 db1
func addData2ToData1(engine1 *xorm.Engine,engine2 *xorm.Engine,ids2 []int64){
	justString := arrayToString(ids2,",")
	sql := fmt.Sprintf("select * from user where user_id in ( %s )", justString )
	var users []User
	err := engine2.SQL(sql).Find(&users)

	println(justString)
	
	//err := engine2.Table("user").Where("user_id in ( ? )", justString ).Find(&users)

	//db2 get data
	 //err := engine2.QueryString(sql).Find(&users)
	if err != nil{
		println(err)
		return
	}
	for i, user := range users{
		println(i,"",user.UserID)
		_, _ = engine1.Insert(&user)
	}
}

func arrayToString(A []int64, delim string) string {
    var buffer bytes.Buffer
    for i := 0; i < len(A); i++ {
        buffer.WriteString( strconv.FormatInt(A[i],10) )
        if i != len(A)-1 {
            buffer.WriteString(delim)
        }
    }
    return buffer.String()
}


// func TestSql111(t *testing.T){
// 	var ids []int64 
// 	ids = append(ids, 12)
// 	ids = append(ids, 22312)
// 	addData2ToData1(nil,nil, ids)
// }