package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type (
  XValidator  struct {
    validator *validator.Validate
  }
  ErrorResponse struct {
    Error       bool
    FailedField string
    Tag         string
    Value       interface{}
  }
)

var Validate = validator.New()

var Valid = XValidator{
    validator: Validate,
}

func (v XValidator) Validate(data interface{}) (error, []string) {
    validationErrors := []ErrorResponse{}

    errs := Validate.Struct(data)
    if errs != nil {
        for _, err := range errs.(validator.ValidationErrors) {
            // In this case data object is actually holding the User struct
            var elem ErrorResponse

            elem.FailedField = err.Field() // Export struct field name
            elem.Tag = err.Tag()           // Export struct tag
            elem.Value = err.Value()       // Export field value
            elem.Error = true

            validationErrors = append(validationErrors, elem)
        }
    }

    if len(validationErrors) > 0 && validationErrors[0].Error {
            errMsgs := make([]string, 0)

            for _, err := range validationErrors {
                errMsgs = append(errMsgs, fmt.Sprintf(
                    "[%s]: '%v' | Needs to implement '%s'",
                    err.FailedField,
                    err.Value,
                    err.Tag,
                ))
            }
            return errs, errMsgs
    }

    return nil, nil 
}

