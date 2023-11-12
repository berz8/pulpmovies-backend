package migrations

import "database/sql"

func CreateMoviesTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS movies (
    	id INTEGER NOT NULL UNIQUE,
    	imdb_id TEXT,
    	original_title TEXT NOT NULL,
    	title TEXT NOT NULL,
    	original_language TEXT,
    	overview TEXT,
    	poster_path TEXT,
    	backdrop_path TEXT,
    	release_date TEXT,
    	runtime INTEGER ,
    	budget INTEGER,
    	revenue INTEGER,
    	adult INTEGER,
    	homepage TEXT,
    	popularity INTEGER,
    	status TEXT,
    	tagline TEXT,
    	video INTEGER,
    	vote_average INTEGER,
    	vote_count INTEGER,
    	deleted_at TEXT,
    	PRIMARY KEY (id) ON CONFLICT FAIL
    );
  `)
	if err != nil {
		return err
	}
	return nil
}
