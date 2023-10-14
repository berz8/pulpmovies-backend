package handlers

import (
	"github.com/berz8/pulpmovies-backend/database"
	"github.com/berz8/pulpmovies-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetUserByID(c *fiber.Ctx) error {
  db := database.DB
  userID := c.Params("id")
  user := models.User{}
  err := db.QueryRow(`
    SELECT * from users WHERE users.id = ?
  `, userID).Scan(
    &user.ID,
    &user.Username,
    &user.Email,
    &user.FullName,
    &user.Birthday,
    &user.Biography,
    &user.ProfilePath,
    &user.AccountStatus,
    &user.Onboarding,
    &user.CreatedAt,
    &user.UpdatedAt,
    &user.DeletedAt,
  )
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "User not found",
    )
  }

  return c.JSON(user)
}

