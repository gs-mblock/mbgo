package jwt

/**
* https://github.com/Wangjiaxing123/JwtDemo
 */

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GenerateToken : 生成令牌
func GenerateToken(c *gin.Context, user CustomClaims, minute int) {
	j := NewJWT("")
	claims := CustomClaims{
		UserID: user.UserID,
		Name:   user.Name,
	}
	claims.NotBefore = int64(time.Now().Unix() - 1000) // 签名生效时间
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()
	//int64(time.Now().Unix() + 3600) // 过期时间 一小时
	claims.Issuer = TokenIssuer //签名的发行者
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Creat token OK",
		"data":    token,
	})
}

// AuthHeader : 中间件，检查token
func AuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(TokenName)
		result := TokenAuth(token)
		if result.Code == 0 && result.Claims != nil {
			c.Set("claims", result.Claims)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    result.Code,
			"message": result.Message,
		})
		c.Abort()
		return
	}
}

// TestGetTwtData 一个需要token认证的测试接口
func TestGetTwtData(c *gin.Context) {
	claims := c.MustGet("claims").(*CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "token有效",
			"data": claims,
		})
	}
}

// TestLogin :test
func TestLogin(c *gin.Context) {
	// check and get userInfo
	user := CustomClaims{
		UserID: 1001,
		Name:   "中国",
	}
	GenerateToken(c, user, 60*24)
}
