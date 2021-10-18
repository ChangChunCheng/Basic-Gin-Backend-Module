package api

import (
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// Get jwt secret key from config
var jwtSecret = []byte(viper.GetString("auth.secret"))

// Get JWT from api.Claims
// return - token string
func GetJWT(claims *Claims) (token string, err error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return token, err
}

// Decode JWT to api.Claims
// return - api.Claims, message for decode error detail
func ParseJWT(auth string) (claims *Claims, message string) {
	token := strings.Split(auth, " ")
	if token[0] != "Bearer" || token[1] == "" {
		return nil, "token not exists"
	}
	tokenClaims, err := jwt.ParseWithClaims(token[1], Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	message = ""
	if err != nil {
		// var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
	}
	claims, ok := tokenClaims.Claims.(*Claims)
	if ok && tokenClaims.Valid {
		// pass
		// fmt.Println("Login Success")
		return claims, ""
	}
	return claims, message
}
