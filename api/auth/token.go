package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken -
func CreateToken(user_id uint32) (string, error) {
  claims := jwt.MapClaims{}
  claims["authorized"] = true
  claims["user_id"] = user_id
  claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expiration
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString([]byte(os.Getenv("API_SECRET")))
}


