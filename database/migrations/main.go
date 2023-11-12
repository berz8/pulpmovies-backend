package migrations

import (
	"database/sql"
)

func RunMigrations(db *sql.DB) error {
	err := CreateUserTable(db)
	if err != nil {
		return err
	}
  err = CreateMoviesTable(db)
	if err != nil {
		return err
	}
	err = CreateWatchlistsTable(db)
	if err != nil {
		return err
	}
  err = CreateWatchlistMoviesTable(db)
	if err != nil {
		return err
	}
	return nil
}
