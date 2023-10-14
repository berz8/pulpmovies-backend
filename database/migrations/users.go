package migrations

import "database/sql"

func CreateUserTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
      id INTEGER NOT NULL UNIQUE,
      username TEXT NOT NULL UNIQUE,
      email TEXT NOT NULL UNIQUE,
      full_name TEXT,
      birthday TEXT,
      biography TEXT,
      profile_path TEXT,
      account_status TEXT NOT NULL,
      onboarding INTEGER NOT NULL,     
      created_at TEXT,
      updated_at TEXT,
      deleted_at TEXT,
      PRIMARY KEY (id) ON CONFLICT FAIL
    )
  `)
	if err != nil {
		return err
	}
	return nil
}
