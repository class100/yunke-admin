package yunke

import (
	`encoding/json`
	"time"

	`github.com/dgrijalva/jwt-go`
	"github.com/storezhang/gox"
)

func token(domain string, method jwt.SigningMethod, secret string) (token string, err error) {
	user := gox.BaseUser{
		BaseStruct: gox.BaseStruct{
			Id: 1,
		},
	}

	// 序列化User对象为JSON
	var userBytes []byte
	if userBytes, err = json.Marshal(user); nil != err {
		return
	}

	claims := jwt.NewWithClaims(method, jwt.StandardClaims{
		// 代表这个JWT的签发主体
		Issuer: domain,
		// 代表这个JWT的主体，即它的所有人
		Subject: string(userBytes),
		// 代表这个JWT的接收对象
		Audience: domain,
		// 是一个时间戳，代表这个JWT的过期时间
		ExpiresAt: time.Now().Add(gox.Month).Unix(),
	})
	token, err = claims.SignedString([]byte(secret))

	return
}
