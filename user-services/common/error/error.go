package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var errValidator = map[string]string{}

func ErrValidationResponse(err error) []ValidationResponse {
	var validationResponses []ValidationResponse
	var fieldErrors validator.ValidationErrors

	if errors.As(err, &fieldErrors) {
		for _, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				validationResponses = append(validationResponses, ValidationResponse{
					Field:   err.Field(),
					Message: "This field is required",
				})
			case "email":
				validationResponses = append(validationResponses, ValidationResponse{
					Field:   err.Field(),
					Message: "Invalid email format",
				})
			default:
				errValidator, ok := errValidator[err.Tag()]
				if ok {
					count := strings.Count(errValidator, "%s")
					if count == 1 {
						validationResponses = append(validationResponses, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field()),
						})
					} else {
						validationResponses = append(validationResponses, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field(), err.Param()),
						})
					}
				} else {
					validationResponses = append(validationResponses, ValidationResponse{
						Field:   err.Field(),
						Message: fmt.Sprintf("something wrong with %s", err.Field()),
					})
				}
			}
		}
	}

	return validationResponses
}

func WrapError(err error)error{
	logrus.Errorf("error: %v", err)
	return err
}
