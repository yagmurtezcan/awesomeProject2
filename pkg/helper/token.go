package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GetToken() (*string, error) {
	var secretKey = []byte("my_secret_key")
	claims := &jwt.MapClaims{
		"iss": "issuer",
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
