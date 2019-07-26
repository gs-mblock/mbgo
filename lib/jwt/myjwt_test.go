package jwt

import (
	"testing"
	"time"
)
func TestJwt(t *testing.T) {
	// creat
	j := NewJWT("")
	claims := CustomClaims{
		UserID: 110,
		Name:   "中国",
	}
	minute := 10;
	claims.NotBefore = int64(time.Now().Unix() - 1000) // 签名生效时间
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()
	claims.Issuer = "Liam" //签名的发行者
	token, err := j.CreateToken(claims)
	if err == nil {
		println(token)
		t.Log("PASS")
	}else{
		println(err)
		t.Fatal("Test_JWT creat jwt failure")
	}
	// check
	result := TokenAuth(token)
	if result.Code == 0 && result.Claims != nil {
		claims2 := result.Claims
		println("claims Name:",claims2.Name)
		t.Log("PASS")
	}else{
		t.Fatal("Test_JWT check jwt failure")
	}
}