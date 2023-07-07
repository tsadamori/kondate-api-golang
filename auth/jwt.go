package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTSignedToken(email, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * 24),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func parseJWTToken(tokenString string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := errors.New("Unexpected signing method")
			return nil, err
		}
		return "", nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}
