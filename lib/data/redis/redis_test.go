package redis

import (
	"fmt"
	"github.com/gs-mblock/mbgo/lib/cache"
	_ "github.com/gs-mblock/mbgo/lib/cache/redis"
	"log"
	"testing"
)

func TestPkgRedis_Cache(t *testing.T) {
	//bm, err := cache.NewCache("redis", `{"conn": "Makeblock123@r-wz9a0dfcef584e14pd.redis.rds.aliyuncs.com:6379"}`)
	bm, err := cache.NewCache("redis", `{"conn": "Makeblock123@localhost:6379"}`)
	if err != nil {
		fmt.Println("err:", err)
		//t.Error("init err")
	}
	log.Printf("ok: %+v\n", bm)
}

func TestPkgRedis_AutoConnect(t *testing.T) {
	sv := new(SvRedis)
	sv.URL = "Makeblock123@localhost:6379"
	sv.AutoConnect()
}

func TestPkgRedis_ConnRedis(t *testing.T) {
	sv := new(SvRedis)
	sv.URL = "Makeblock123@localhost:6379"
	v := sv.NewRedisCache()
	println("result:", v)
}

func TestPKG_Redis_data(t *testing.T) {
	sv := new(SvRedis)
	sv.URL = "Makeblock123@localhost:6379"
	sv.NewRedisCache()
	n := "parent-1|heZi_u1"
	f := sv.IsExist(n)
	log.Println("IsExist=", f)
	sv.Set(n, "xxxx-xxx", 100)
	v, _ := sv.Get(n)
	log.Println("v=", v)
	f2 := sv.IsExist(n)
	log.Println("IsExist=", f2)
	sv.Delete(n)
	f3 := sv.IsExist(n)
	log.Println("IsExist=", f3)
}

func TestPkg_RedisByte(t *testing.T) {
	sv := new(SvRedis)
	sv.URL = "Makeblock123@localhost:6379"
	sv.NewRedisCache()
	n := "parent-1|heZi_u1"
	sv.Set(n, []byte("123456-xxx"), 100)
	v, err := sv.GetBytes(n)
	log.Printf("err=:%+v\n", err)
	log.Printf("v=:%+v\n", string(v))
}
