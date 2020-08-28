package util

import (
	setting "ginExample/pkg"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (token string, err error) {
	// 获取当前时间
	nowTime := time.Now()
	// 设置超时时间
	expireTime := nowTime.Add(3 * time.Hour)

	// 定义claim
	claims := Claims{
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)

}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claim, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claim, nil
		}
	}
	return nil, err
}
