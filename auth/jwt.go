package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO
var secretKey string = "SECRET_KEY"
var TokenKey string = "token"

func CreateJWTSignedToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWTToken(tokenString string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := errors.New("Unexpected signing method")
			return nil, err
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}
