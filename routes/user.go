package routes

import (
	"github.com/berz8/pulpmovies-backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
  user := app.Group("/user")

  user.Get("/:id", handlers.GetUserByID)


}
