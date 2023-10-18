package validators

type AuthGoogle struct {
  IdToken string `json:"idToken" validate:"required"`
  FullName string `json:"fullName" validate:"required"`
}
