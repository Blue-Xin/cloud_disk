package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"zero/greet/define"
)

func Md5(s string) string {
	fmt.Println(md5.Sum([]byte(s)))
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func GenerateToken(id int, identity, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, uc)
	signedString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}
