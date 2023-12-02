package handlers

import (
	"context"
	"database/sql"
	"strings"

	"github.com/berz8/pulpmovies-backend/database"
	"github.com/berz8/pulpmovies-backend/models"
	"github.com/berz8/pulpmovies-backend/validators"
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

func GetUserByUsername(c *fiber.Ctx) error {
  db := database.DB
  userID := c.Params("username")
  user := models.User{}
  err := models.GetUserByUsername(&user, db, userID)
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "User not found" + err.Error(),
    )
  }

  return c.JSON(user)
}

func CheckUsername(c *fiber.Ctx) error {
  db := database.DB
  userID := c.Params("username")
  if len(userID) < 4 {
    return fiber.NewError(
      fiber.StatusNotAcceptable,
      "Username must be 4 characters",
    )
  }
  user := models.User{}
  usernameExist := true
  err := models.GetUserByUsername(&user, db, userID)
  if err != nil {
    usernameExist = false 
  }

  return c.JSON(fiber.Map{ "result": usernameExist })
}

func OnBoarding(c *fiber.Ctx) error {
  onBoardingBody := new(validators.OnBoard)
  if err := c.BodyParser(onBoardingBody); err != nil {
    return err
  }

  err, errMsgs := validators.Valid.Validate(onBoardingBody)
  if err != nil {
    return &fiber.Error{
      Code:    fiber.StatusBadRequest,
      Message: strings.Join(errMsgs, " and "),
    }
  }
  db := database.DB
  user := models.User{}
  err = models.GetUserByUsername(&user, db, onBoardingBody.Username,)
  if err != nil {
    if err != sql.ErrNoRows {
      return &fiber.Error{
        Code:    fiber.StatusBadRequest,
        Message: "Something went wrong " + err.Error(),
      }
    }
  } else if user.Username == onBoardingBody.Username {
    return &fiber.Error{
      Code:    fiber.StatusBadRequest,
      Message: "Username already in use",
    }
  }


  userID, _ := models.GetUserIDFromToken(c)

  _, err = db.ExecContext(
    context.Background(),
    `UPDATE users SET username = ?, onboarding = 1 WHERE id = ?`,
    onBoardingBody.Username,
    userID,
  )
  if err != nil {
    return fiber.NewError(
      fiber.StatusBadRequest,
      "Something went wrong while setting username" + err.Error(),
    )
  }

  return c.JSON(fiber.Map{ "message": "User onboarded successfully" }) 

}
