package routes

import (
	"os"

	"github.com/berz8/pulpmovies-backend/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)


func UserRoutes(app *fiber.App) {
  accessSecret := os.Getenv("ACCESS_SECRET")
  user := app.Group("/user")

  user.Get("/id/:id", handlers.GetUserByID)
  user.Get("/username/:id", handlers.GetUserByUsername)
  user.Get("/username/:username/check", handlers.CheckUsername)
  user.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(accessSecret)},
  }))
  user.Post("/onboarding", handlers.OnBoarding)


}
