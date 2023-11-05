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
  err := models.GetUserByID(&user, db, userID)
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "User not found",
    )
  }

  return c.JSON(user)
}

