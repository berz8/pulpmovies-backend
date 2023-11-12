package migrations

import "database/sql"

func CreateWatchlistsTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS watchlists (
      id INTEGER NOT NULL UNIQUE,
      name TEXT NOT NULL,
      description TEXT,
      user_id INTEGER NOT NULL,     
      public INTEGER NOT NULL,     
      PRIMARY KEY (id) ON CONFLICT FAIL
      FOREIGN KEY (user_id) REFERENCES users(id)
    )
  `)
	if err != nil {
		return err
	}
	return nil
}

func CreateWatchlistMoviesTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS watchlist_movies (
      watchlist_id INTEGER NOT NULL UNIQUE,
      movie_id INTEGER NOT NULL UNIQUE,
      PRIMARY KEY (watchlist_id, movie_id) ON CONFLICT FAIL
      FOREIGN KEY (watchlist_id) REFERENCES watchlists(id)
      FOREIGN KEY (movie_id) REFERENCES movies(id)
    )
  `)
	if err != nil {
		return err
	}
	return nil
}
