package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
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

// 邮箱验证码发送
func MailSendCode(mail, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "1251780434@qq.com")
	m.SetHeader("To", mail)
	m.SetHeader("Subject", "BlueXin!")
	m.SetBody("text/html", "验证码： <h1>"+code+"</h1>")

	d := gomail.NewDialer("smtp.qq.com", 587, "1251780434@qq.com", define.MailPassword)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// 获取随机数
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
