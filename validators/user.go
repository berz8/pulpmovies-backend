package validators

type OnBoard struct {
  Username string `json:"username" validate:"required,min=4,max=30"`
}
