package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/the7s/go-vue-blog/server/model/system/request"
)

type JWT struct {
	SigningKey []byte
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(j.SigningKey)
	return token.SignedString(j.SigningKey)
}
