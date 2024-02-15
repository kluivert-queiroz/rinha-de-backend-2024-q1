package validator

import "github.com/go-playground/validator/v10"

type (
	XValidator struct {
		validator *validator.Validate
	}
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
func (x XValidator) Validator() *validator.Validate {
	return x.validator
}

func NewValidator() *XValidator {
	return &XValidator{
		validator: validate,
	}
}
