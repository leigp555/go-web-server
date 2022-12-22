package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go/note/config"
	"time"
)

type MyCustomClaims struct {
	User string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (tokenStr string, err error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(config.GlobalConfig.Token.ExpiresTime), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",     //颁发者
			Subject:   "somebody", //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	tokenStr, err = token.SignedString([]byte(config.GlobalConfig.Token.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenStr, err
}

func ParseToken(tokenStr string) (userId string, err error) {
	fmt.Println(tokenStr)
	token, _ := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalConfig.Token.SigningKey), nil
	})
	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		return claims.User, nil
	}
	err = errors.New("用户为认证")
	return "", err

}
