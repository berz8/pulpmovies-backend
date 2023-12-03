package handlers

import (
	"strings"

	"github.com/berz8/pulpmovies-backend/database"
	"github.com/berz8/pulpmovies-backend/models"
	"github.com/berz8/pulpmovies-backend/validators"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
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

func AddMovieToWatchlist(c *fiber.Ctx) error {
  db := database.DB
  watchlistID := c.Params("id")
  movieBody := new(validators.Movie)

  if err := c.BodyParser(movieBody); err != nil {
    return err
  }

  err, errMsgs := validators.Valid.Validate(movieBody)
  if err != nil {
    return &fiber.Error{
      Code:    fiber.ErrBadRequest.Code,
      Message: strings.Join(errMsgs, " and "),
    }
  }

  userID, err := models.GetUserIDFromToken(c)
  if err != nil {
    return fiber.NewError(
      fiber.StatusBadRequest,
      "Something went wrong while getting user id" + err.Error(),
    )
  }

  watchlist := models.Watchlist{}
  if watchlistID == "default" {
    err = models.GetDefaultWatchlistByUserID(&watchlist, db, int32(userID))
  } else {
    err = models.GetWatchlistByID(&watchlist, db, watchlistID)
  }
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "Watchlist not found",
    )
  }


  if watchlist.UserID != int32(userID) {
    return fiber.NewError(
      fiber.StatusBadRequest,
      "Watchlist does not belong to user",
    )
  }

  movie := models.Movie{
    ID: movieBody.ID,
    ImdbID: null.NewString(movieBody.ImdbID, true),
    OriginalTitle: movieBody.OriginalTitle,
    Title: movieBody.Title,
    OriginalLanguage: null.NewString(movieBody.OriginalLanguage, true),
    Overview: null.NewString(movieBody.Overview, true),
    PosterPath: null.NewString(movieBody.PosterPath, true),
    BackdropPath: null.NewString(movieBody.BackdropPath, true),
    ReleaseDate: null.NewString(movieBody.ReleaseDate, true),
  }

  err = models.CreateMovie(&movie, db)
  if err != nil {
    return fiber.NewError(
      fiber.StatusBadRequest,
      "Something went wrong while adding movie" + err.Error(),
    )
  }

  err = models.AddMovieToWatchlist(db, watchlist.ID, movie.ID)
  if err != nil {
    if err.Error() == "UNIQUE constraint failed: watchlist_movies.movie_id, watchlist_movies.watchlist_id" {
      return c.SendStatus(fiber.StatusOK)
    }
    return fiber.NewError(
      fiber.StatusBadRequest,
      "Something went wrong while adding movie to watchlist" + err.Error(),
    )
  }

  return c.SendStatus(fiber.StatusCreated)
}

func GetIsMovieInUserWatchlists(c *fiber.Ctx) error {
  db := database.DB
  movieID := c.Params("movieId")

  userID, err := models.GetUserIDFromToken(c)
  if err != nil {
    return fiber.NewError(
      fiber.StatusBadRequest,
      "Something went wrong while getting user id" + err.Error(),
    )
  }

  watchlists := []models.Watchlist{}
  err = models.UserWatchlistsHaveMovie(db, int32(userID), movieID, &watchlists)
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "Watchlist not found",
    )
  }

  return c.JSON(watchlists)
}
