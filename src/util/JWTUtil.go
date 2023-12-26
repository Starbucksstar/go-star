package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"star/src/global"
	"time"
)

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Id       string `json:"id"`
	jwt.RegisteredClaims
}

func GenToken(username, role, id string) (string, error) {
	claims := CustomClaims{
		username,
		role,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(global.TokenExpireDuration)),
			Issuer:    "star-go",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(global.CustomSecret)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return global.CustomSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
