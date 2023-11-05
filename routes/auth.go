package routes

import (
	"github.com/berz8/pulpmovies-backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
  auth := app.Group("/auth")

  auth.Post("/google", handlers.AuthGoogle)


}
