package xorm

import (
	//"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	//log "github.com/skoo87/log4go"
	//"os"
	//"path/filepath"
)

// DBEngine mysql
// var DBEngine *xorm.Engine

// ConnectXorm mysql: 
/**
"root:1234qwer@tcp(127.0.0.1:3306)/msdb" 
 "t_"
*/
func ConnectXorm(url string, prefix string) (*xorm.Engine , error){
	DBEngine, err := xorm.NewEngine("mysql", url)
	tbMapper := core.NewPrefixMapper(core.GonicMapper{}, prefix)
	DBEngine.SetTableMapper(tbMapper)
	DBEngine.SetColumnMapper(core.GonicMapper{})
	if err != nil {
		println("conf init error:%s", err.Error())
		return DBEngine, err
	}
	err = DBEngine.Ping()
	if err != nil {
		println("conf init error:%s", err.Error())
		return DBEngine, err
	}
	return DBEngine, nil
}