package models

import (
	"database/sql"

	"gopkg.in/guregu/null.v4"
)

type User struct {
  ID int32 `json:"id"`
  Username string `json:"username"`
  Email string `json:"email"`
  FullName null.String `json:"fullName"`
  Birthday null.String `json:"birthday"`
  Biography null.String `json:"biography"`
  ProfilePath null.String `json:"profilePath"`
  AccountStatus string `json:"accountStatus"`
  Onboarding int8 `json:"onboarding"`
  CreatedAt null.String `json:"-"`
  UpdatedAt null.String `json:"-"`
  DeletedAt null.String `json:"-"`
}

func GetUserByID(user *User, db *sql.DB, id string) error {
 err := db.QueryRow(`
    SELECT * from users WHERE users.id = ?
  `, id).Scan(
    &user.ID,
    &user.Username,
    &user.Email,
    &user.FullName,
    &user.Birthday,
    &user.Biography,
    &user.ProfilePath,
    &user.AccountStatus,
    &user.Onboarding,
    &user.CreatedAt,
    &user.UpdatedAt,
    &user.DeletedAt,
  )
  return err
}

func GetUserByEmail(user *User, db *sql.DB, email string) error {
 err := db.QueryRow(`
    SELECT * from users WHERE users.email = ?
  `, email).Scan(
    &user.ID,
    &user.Username,
    &user.Email,
    &user.FullName,
    &user.Birthday,
    &user.Biography,
    &user.ProfilePath,
    &user.AccountStatus,
    &user.Onboarding,
    &user.CreatedAt,
    &user.UpdatedAt,
    &user.DeletedAt,
  )
  return err
}

func GetUserByUsername(user *User, db *sql.DB, username string) error {
 err := db.QueryRow(`
    SELECT * from users WHERE users.username = ?
  `, username).Scan(
    &user.ID,
    &user.Username,
    &user.Email,
    &user.FullName,
    &user.Birthday,
    &user.Biography,
    &user.ProfilePath,
    &user.AccountStatus,
    &user.Onboarding,
    &user.CreatedAt,
    &user.UpdatedAt,
    &user.DeletedAt,
  )
  return err
}
