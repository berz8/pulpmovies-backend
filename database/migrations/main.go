package migrations

import (
	"database/sql"
)

func RunMigrations(db *sql.DB) error {
  err := CreateUserTable(db)

  if err != nil {
    return err
  }
  return nil
}
