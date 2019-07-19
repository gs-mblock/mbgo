package jwt

import (
	"errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// 一些常量
const (
	// SignKey ：加密盐,每段时间更新盐:新旧盐可同时进行
	SignKey = "188274e3052a9fd6450077a3287003a-mb1907@jm@!$$v2"
	// SignKeyOld 上一个加密盐
	SignKeyOld = "188274e3052a9fd6450077a3287003a3-mb1907@jm@!$$"
	// SignKeyExpiredTime : 上个盐有效时间, 1491888244, 1564650738
	SignKeyExpiredTime int64 = 1491888244
	// TokenName header name
	TokenName   = "utoken"
	TokenIssuer = "makeblock"
)

var (
	// ErrorTokenExpired : 过期了
	ErrorTokenExpired = errors.New("Token is expired")
	// ErrorTokenNotValidYet ：还没有生效
	ErrorTokenNotValidYet = errors.New("Token not active yet")
	// ErrorTokenMalformed  ：That's not even a token
	ErrorTokenMalformed = errors.New("That's not even a token")
	// ErrorTokenInvalid ：Couldn't handle this token
	ErrorTokenInvalid = errors.New("Couldn't handle this token")
)

// AuthResult : back result
type AuthResult struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Claims  *CustomClaims
}

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserID int64  `json:"id"`
	Name   string `json:"name"`
	//Email  string `json:"email"`
	//Phone  string `json:"phone"`
	//Role   string `json:"role"` //角色
	//Info   string `json:"info"`
	jwtgo.StandardClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// TokenAuth :JWTAuth 中间件，检查token
func TokenAuth(token string) AuthResult {
	var result AuthResult
	if token == "" {
		result.Code = 4101
		result.Message = "请求未携带token，无权限访问"
		return result
	}
	j := NewJWT(SignKey)
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == ErrorTokenExpired {
			result.Code = 4102
			result.Message = "授权已过期"
			return result
		}
		// 查看旧的盐
		claims2, err2 := ParseOldToken(token)
		if err2 == nil {
			result.Code = 0
			result.Message = "用过期盐授权"
			result.Claims = claims2
			return result
		}
		result.Code = 4103
		result.Message = err.Error()
		return result
	}
	result.Code = 0
	result.Message = "OK"
	result.Claims = claims
	return result
	// 继续交由下一个路由处理,并将解析出的信息传递下去
	//c.Set("claims", claims)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, ErrorTokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrorTokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, ErrorTokenNotValidYet
			} else {
				return nil, ErrorTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrorTokenInvalid
}

// ParseOldToken :check old tocken
func ParseOldToken(tokenString string) (*CustomClaims, error) {
	//println("old-time %d: ", time.Now().Unix())
	if SignKeyExpiredTime <= time.Now().Unix() {
		return nil, ErrorTokenInvalid
	}
	j := NewJWT(SignKeyOld)
	println("Use old SignKey ")
	return j.ParseToken(tokenString)
}

// RefreshToken 更新token,默认加 1小时
func (j *JWT) RefreshToken(tokenString string, minute int) (string, error) {
	jwtgo.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwtgo.TimeFunc = time.Now
		if minute <= 0 {
			claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		} else {
			claims.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()
		}
		return j.CreateToken(*claims)
	}
	return "", ErrorTokenInvalid
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// NewJWT 新建一个jwt实例
func NewJWT(key string) *JWT {
	if key == "" {
		key = SignKey
	}
	return &JWT{
		[]byte(key),
	}
}
