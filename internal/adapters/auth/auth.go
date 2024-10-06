package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NikolayStrekalov/practicum-gophermart/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

const TOKEN_EXP = time.Hour * 24
const SECRET_KEY = "supersecretkey" // to be moved to secret store

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
}

func BuildJWTString(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
		UserID: id,
	})
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func SetAuthorization(res *http.ResponseWriter, user *models.User) error {
	token, err := BuildJWTString(user.ID)
	if err != nil {
		return fmt.Errorf("error creating token: %w", err)
	}
	(*res).Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return nil
}
