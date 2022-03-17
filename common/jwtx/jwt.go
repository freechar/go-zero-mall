package jwtx

import (

	"github.com/golang-jwt/jwt"
)
func GetToken(secretKey []byte,iat,seconds,uid int64)(string,error) {
	claims:=make(jwt.MapClaims)
	claims["exp"] =iat+seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	// token :=jwt.New(jwt.SigningMethodES256)
	// token.Claims = claims
	return token.SignedString(secretKey)
}
