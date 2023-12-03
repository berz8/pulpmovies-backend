package models

import (
	"database/sql"
	"reflect"

	"gopkg.in/guregu/null.v4"
)

type Watchlist struct {
  ID int32 `json:"id"`
  Name string `json:"name"`
  Description null.String `json:"description"`
  Public int8 `json:"public"`
  IsDefault int8 `json:"isDefault"`
  UserID int32 `json:"userId"`
}

type WatchlistMovie struct {
  ID int32 `json:"id"`
  OriginalTitle string `json:"original_title"`
  Title string `json:"title"`
  PosterPath null.String `json:"poster_path"`
  ReleaseDate null.String `json:"release_date"`
}



func GetWatchlistByID(watchlist *Watchlist, db *sql.DB, id string) error {
 err := db.QueryRow(`
    SELECT * from watchlists WHERE watchlists.id = ?
  `, id).Scan(
    &watchlist.ID,
    &watchlist.Name,
    &watchlist.Description,
    &watchlist.Public,
    &watchlist.IsDefault,
    &watchlist.UserID,
  )
  return err
}

func GetWatchlistByUserID(watchlists *[]Watchlist, db *sql.DB, id string) error {
  rows, err := db.Query(`
    SELECT * from watchlists WHERE watchlists.user_id = ?
  `, id) 

  if err != nil {
    return err
  }

  destv := reflect.ValueOf(watchlists).Elem()

    args := make([]interface{}, destv.Type().Elem().NumField())

    for rows.Next() {

        rowp := reflect.New(destv.Type().Elem())
        rowv := rowp.Elem()

        for i := 0; i < rowv.NumField(); i++ {
            args[i] = rowv.Field(i).Addr().Interface()
        }

        if err := rows.Scan(args...); err != nil {
            return err
        }

        destv.Set(reflect.Append(destv, rowv))
      }

  return err
}

func GetDefaultWatchlistByUserID(watchlist *Watchlist, db *sql.DB, userID int32) error {
  err := db.QueryRow(`
  SELECT * from watchlists WHERE watchlists.user_id = ? AND watchlists.is_default = 1
  `, userID).Scan(
    &watchlist.ID,
    &watchlist.Name,
    &watchlist.Description,
    &watchlist.Public,
    &watchlist.IsDefault,
    &watchlist.UserID,
  )
  return err
}

func GetWatchlistMovies(movies *[]WatchlistMovie, db *sql.DB, id string) error {
  rows, err := db.Query(`
    SELECT 
      m.id as id, 
      m.original_title as original_title,
      m.title as title,
      m.poster_path as poster_path,
      m.release_date as release_date
      FROM watchlist_movies wm
      LEFT JOIN movies m ON wm.movie_id = m.id
      WHERE wm.watchlist_id = ?
  `, id) 

  if err != nil {
    return err
  }

  destv := reflect.ValueOf(movies).Elem()

    args := make([]interface{}, destv.Type().Elem().NumField())

    for rows.Next() {

        rowp := reflect.New(destv.Type().Elem())
        rowv := rowp.Elem()

        for i := 0; i < rowv.NumField(); i++ {
            args[i] = rowv.Field(i).Addr().Interface()
        }

        if err := rows.Scan(args...); err != nil {
            return err
        }

        destv.Set(reflect.Append(destv, rowv))
      }

  return err
}

func AddMovieToWatchlist(db *sql.DB, watchlistID int32, movieID int32) error {
  _, err := db.Exec(`
    INSERT INTO watchlist_movies (watchlist_id, movie_id) VALUES (?, ?)
  `, watchlistID, movieID)
  return err
}

func UserWatchlistsHaveMovie(db *sql.DB, userID int32, movieID string, watchlists *[]Watchlist) error {
  rows, err := db.Query(`
    SELECT w.id as id, w.name as name, w.description as description, w.public as public, w.is_default as is_default, w.user_id as user_id
    FROM watchlist_movies wm 
    LEFT JOIN watchlists w ON wm.watchlist_id = w.id
    WHERE wm.movie_id = ? AND w.user_id = ?
  `, movieID, userID)
  if err != nil {
    return err
  }

  for rows.Next() {
    watchlist := Watchlist{}
    err = rows.Scan(
      &watchlist.ID,
      &watchlist.Name,
      &watchlist.Description,
      &watchlist.Public,
      &watchlist.IsDefault,
      &watchlist.UserID,
    )
    if err != nil {
      return err
    }
    *watchlists = append(*watchlists, watchlist)
  }

  return err
}


func RemoveMovieFromWatchlist(db *sql.DB, watchlistID int32, movieID string) error {
  _, err := db.Exec(`
    DELETE FROM watchlist_movies WHERE watchlist_id = ? AND movie_id = ?
  `, watchlistID, movieID)
  return err
}


