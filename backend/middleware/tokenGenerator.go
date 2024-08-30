package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID string) (string, error) {
	//Token oluştur
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	//Token imza
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
