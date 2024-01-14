package common

import (
	"openapi/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signKey = []byte(global.Config.Jwt.SignKey)

// GeneratorToken 生成JWT token
func GeneratorToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return ""
	}
	return tokenString
}

// VerifyToken 验证token
func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", nil
		}
		return username, nil
	}
	return "", &tokenNotValidate
}
