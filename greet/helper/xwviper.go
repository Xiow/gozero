package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

)

//定义一些过期时间


// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("夏天夏天悄悄过去")
const TokenExpiredTime=time.Hour*24
type MyClaims struct {
	Name string 	`json:"name"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}
func GenerateToken(name,password string)(string,error){
	mc:=MyClaims{
		name,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiredTime)),
			Issuer:    "xwg", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}


