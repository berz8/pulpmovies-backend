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
