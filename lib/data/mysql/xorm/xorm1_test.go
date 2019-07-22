package xorm

import (
	"testing"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var DBEngine *xorm.Engine

func TestXorm(t *testing.T){
	//dbURL := "root:1234qwer@tcp(127.0.0.1:3306)/msdb"
	dbURL := "root:1234qwer@(127.0.0.1:3306)/test?charset=utf8"
	println("[Test]",dbURL)
	DBEngine, err := xorm.NewEngine("mysql", dbURL)
	println("[Test]-1")
	tbMapper := core.NewPrefixMapper(core.GonicMapper{}, "t_")
	DBEngine.SetTableMapper(tbMapper)
	DBEngine.SetColumnMapper(core.GonicMapper{})
	if err != nil {
		println("conf init error:%s", err.Error())
	}
	err = DBEngine.Ping()
	if err != nil {
		println("conf init error:%s", err.Error())
	}
}


func TestConnect(t *testing.T) {
	url := "root:1234qwer@tcp(127.0.0.1:3306)/test"
	println(url)
	DBEngine, err := xorm.NewEngine("mysql", url)
	println("ok-1")
	// tbMapper := core.NewPrefixMapper(core.GonicMapper{}, "t_")
	// DBEngine.SetTableMapper(tbMapper)
	// DBEngine.SetColumnMapper(core.GonicMapper{})

	//DBEngine ,err := ConnectXorm(url,"cp_")
	println(err)
	if err !=nil{
		t.Fatal(err)
	}
	println(DBEngine)

	// creat table
	// sql :=`DROP TABLE IF EXISTS test.tb_test; CREATE TABLE test.tb_test (
	// 	dbid BIGINT NOT NULL,
	// 	name VARCHAR(45) NULL,
	// 	PRIMARY KEY (dbid));
	// 	`
	// result, err :=DBEngine.Exec(sql)
	// if err !=nil {
	// 	t.Fatal(err)
	// }
	// println(result)


}