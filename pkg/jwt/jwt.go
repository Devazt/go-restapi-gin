package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "keyboard cat"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return webtoken, nil
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := ValidateToken(tokenStr)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
