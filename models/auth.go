package models

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userID int64) (string, error) {

  tokenClaims := jwt.MapClaims{}
  tokenClaims["user_id"] = userID
  tokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
  tokenClaims["authorized"] = true

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
  tokenSigned, err := token.SignedString([]byte(os.Getenv("ACCES_SECRET")))
  
  return tokenSigned, err
}
