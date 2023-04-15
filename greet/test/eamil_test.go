package test

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"testing"
	"zero/greet/define"
)

func TestEmail(t *testing.T) {
	m := gomail.NewMessage()
	m.SetHeader("From", "1251780434@qq.com")
	m.SetHeader("To", "hashmap0427@gmail.com")
	m.SetHeader("Subject", "BlueXin!")
	m.SetBody("text/html", "验证码： <h1>123456</h1>")

	d := gomail.NewDialer("smtp.qq.com", 587, "1251780434@qq.com", define.MailPassword)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
