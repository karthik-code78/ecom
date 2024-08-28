package auth

import (
	"github.com/karthik-code78/ecom/shared/configure"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(configure.GetJwtSecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
