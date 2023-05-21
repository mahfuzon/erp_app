package helper

import "github.com/go-playground/validator/v10"

func FormatErrorValidation(errors validator.ValidationErrors) []string {
	var errorMessages []string
	for _, fieldError := range errors {
		errorMessages = append(errorMessages, fieldError.Error())
	}

	return errorMessages
}
