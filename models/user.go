package models

import (
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
