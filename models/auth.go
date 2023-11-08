package models

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID int64) (string, error) {

  tokenClaims := jwt.MapClaims{}
  tokenClaims["user_id"] = userID
  tokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
  tokenClaims["authorized"] = true

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
  tokenSigned, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  
  return tokenSigned, err
}


func GetUserIDFromToken(c *fiber.Ctx) (float64, error) {
  token := c.Locals("user").(*jwt.Token)
  claims := token.Claims.(jwt.MapClaims)
  userID := claims["user_id"].(float64)

  return userID, nil
}
