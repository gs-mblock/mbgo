package xorm
/**
https://blog.csdn.net/yiweiyi329/article/details/90442484 
https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/56028
*/
import (
	"testing"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func TestXorm2(t *testing.T){
	url := "root:1234qwer@(localhost:3306)/test?charset=utf8"
    engine, err := xorm.NewEngine("mysql", url)
    if err != nil{
		fmt.Println(err)
		t.Fatal(err)
	}
	println(engine)
	//creatTable(engine)
	insert(engine)
	delete(engine)
	
}

func creatTable(engine *xorm.Engine){
	//建表
    err := engine.Sync2(new(User))
    if err != nil{
        fmt.Println("error in create table user, ", err)
	}
}

type User struct {
    ID int64
    Name string
    Age int
    Passwd string `xorm:"varchar(200)"`
    Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Salt string
}

func insert(engine *xorm.Engine){
	var user User
	user.Name = "guigui"
	user.Salt = "hell"
	user.Age = 18
	user.Passwd = "23"
	affected, err := engine.Insert(&user)
	if err != nil{
		fmt.Println(err)
	}
	println("[Test]", affected)
}

func delete(engine *xorm.Engine){
	var user User
	affected, err := engine.ID(2).Delete(&user)
	if err != nil{
		fmt.Println(err)
	}
	println("[Test]", affected)
}

func edit(engine *xorm.Engine){
	mm, err := engine.Exec("update user set name = ? where id = ?", "lgr", 1)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(mm)
}

func queryString(engine *xorm.Engine){
	results, err := engine.QueryString("select * from user")
	if err != nil{
		fmt.Println(err)
	}
	for k, v := range results{
		fmt.Println(k)
		for m, n := range v{
			fmt.Println(m)
			fmt.Println(n)
		}
		fmt.Println("=====")
	}
}

func query(engine *xorm.Engine){
	var users []User
	err := engine.Table("user").Where("name = ?", "lgr").And("age > 10").Find(&users)
	//以上相当于sql语句：SELECT * FROM user WHERE name = "lgr" AND age > 10
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(users)
}

//实现事务的两种方法：


