package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

// function to generate jwt token
func GenerateJWT(payload map[string]string) (string, error) {

	if payload == nil {
		return "", fmt.Errorf("error in generating token: payload is nil")
	}

	claims := jwt.MapClaims{
		"payload": payload,
		"exp":	 jwt.TimeFunc().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // creating a new token with the claims and the signing method HS256

	tokenString, err := token.SignedString(secretKey) // signing the token with the secret key
	if err != nil {
		return "", fmt.Errorf("error in generating token: %v", err)
	}

	return tokenString, nil
}


// function to validate jwt token
func ValidateJWT(tokenString string) (*jwt.Token, error) { 

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error in parsing the token")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error in validating token: %v", err)
	}

	return token ,nil

}