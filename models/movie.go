package models

import (
	"database/sql"

	"gopkg.in/guregu/null.v4"
)

type Movie struct {
  ID int32 `json:"id"`
  ImdbID null.String `json:"imdbId"`
  OriginalTitle string `json:"originalTitle"`
  Title string `json:"title"`
  OriginalLanguage null.String `json:"originalLanguage"`
  Overview null.String `json:"overview"`
  PosterPath null.String `json:"posterPath"`
  BackdropPath null.String `json:"backdropPath"`
  ReleaseDate null.String `json:"releaseDate"`
}




func GetMovieByID(movie *Movie, db *sql.DB, id string) error {
 err := db.QueryRow(`
    SELECT * from movie WHERE movie.id = ?
  `, id).Scan(
    &movie.ID,
    &movie.ImdbID,
    &movie.OriginalTitle,
    &movie.Title,
    &movie.OriginalLanguage,
    &movie.Overview,
    &movie.PosterPath,
    &movie.ReleaseDate,
  )
  return err
}
