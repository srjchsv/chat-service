package services

import (
	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "your-secret-key"

func ValidateToken(tokenString string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		username := claims["username"].(string)
		return userID, username, nil
	}

	return 0, "", err
}
