package auth

import (
	"github.com/Yummiii/Nayu/internal/database"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *database.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
	})

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
