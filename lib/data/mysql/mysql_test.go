package mysql

import (
	"testing"
)

func TestMySQL_Connect(t *testing.T) {
	//initMySQLData(engine, configs.EnvConfig.MysqlDBName)
	//initMySQLData(EngineOrg, configs.EnvConfig.MysqlDBNameOrg)
	sv := new(SvMySql)
	sv.URL = "root:1234qwer@tcp(127.0.0.1:3306)/test?charset=utf8mb4" //"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4"
	sv.NewEngine()
	println("statue:", sv.IsConnect)
	//log.Printf("v:%+v\n",sv.engine)
}
