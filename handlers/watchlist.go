package handlers

import (
	"github.com/berz8/pulpmovies-backend/database"
	"github.com/berz8/pulpmovies-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetWatchlistByID(c *fiber.Ctx) error {
  db := database.DB
  watchlistID := c.Params("id")
  watchlist := models.Watchlist{}
  err := models.GetWatchlistByID(&watchlist, db, watchlistID)
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "Watchlist not found",
    )
  }
  movies := []models.WatchlistMovie{}
  err = models.GetWatchlistMovies(&movies, db, watchlistID)
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "Watchlist not found",
    )
  }
  return c.JSON(fiber.Map{
    "watchlist": watchlist,
    "movies": movies,
  })
}

func GetWatchlistByUserID(c *fiber.Ctx) error {
  db := database.DB
  watchlistID := c.Params("id")
  watchlist := []models.Watchlist{}
  err := models.GetWatchlistByUserID(&watchlist, db, watchlistID)
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "User not found",
    )
  }
  return c.JSON(watchlist)
}
