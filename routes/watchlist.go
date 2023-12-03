package routes

import (
	"os"

	"github.com/berz8/pulpmovies-backend/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)


func WatclistRoutes(app *fiber.App) {
  accessSecret := os.Getenv("ACCESS_SECRET")
  user := app.Group("/watchlist")

  user.Get("/id/:id", handlers.GetWatchlistByID)
  user.Get("/user/:id", handlers.GetWatchlistByUserID)
  user.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(accessSecret)},
  }))
  user.Post("/id/:id/add", handlers.AddMovieToWatchlist)
  user.Delete("/id/:id/:movieId", handlers.RemoveMovieFromWatchlist)
  user.Get("/user/movie/:movieId", handlers.GetIsMovieInUserWatchlists)


}
