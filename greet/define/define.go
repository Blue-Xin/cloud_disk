package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var MailPassword = "oduugmggonfjifeb"
var JwtKey = "cloud-disk-key"

//验证码长度

var CodeLength = 6

// 设置过期时间（s）

var CodeExpire = 300
