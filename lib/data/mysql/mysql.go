package mysql

import (
	_ "github.com/go-sql-driver/mysql" //go-lint
	"github.com/go-xorm/xorm"
	"log"
	"time"
	"xorm.io/core"
)

type SvMySql struct {
	URL         string //
	IsConnect   bool
	connectTime int64
	engine      *xorm.Engine
}

func (sv *SvMySql) Engine() *xorm.Engine {
	if !sv.IsConnect {
		sv.AutoConnect()
	}
	return sv.engine
}

// AutoConnect :
func (sv *SvMySql) AutoConnect() {
	if !sv.IsConnect {
		return
	}
	now := time.Now().Unix()
	// 10*60 秒
	if now-sv.connectTime < 600 {
		return
	}
	sv.NewEngine()
}

// NewEngine :
func (sv *SvMySql) NewEngine() bool {
	sv.connectTime = time.Now().Unix()
	var err error
	//MySQLMasterURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", configs.EnvConfig.MysqlUser, configs.EnvConfig.MysqlPwd, configs.EnvConfig.MysqlHost, configs.EnvConfig.MysqlPort, dbName)
	log.Println("[INFO] MySQLMasterURL 1=", sv.URL)
	sv.engine, err = xorm.NewEngine("mysql", sv.URL)
	tbMapper := core.NewPrefixMapper(core.GonicMapper{}, "")
	sv.engine.SetTableMapper(tbMapper)
	sv.engine.SetColumnMapper(core.GonicMapper{})
	//engine.ShowSQL(true)
	if err != nil {
		log.Printf("[ERROR]mysql init error:%s", err.Error())
		sv.IsConnect = false
		return false
	}
	err = sv.engine.Ping()
	if err != nil {
		log.Printf("[ERROR]mysql ping error:%s", err.Error())
		sv.IsConnect = false
		return false
	}
	sv.engine.SetMaxIdleConns(100)
	sv.engine.SetMaxOpenConns(500)
	sv.engine.SetConnMaxLifetime(100 * time.Second)
	sv.IsConnect = true
	log.Println("[INFO] mysql connect ok")
	return true
}
