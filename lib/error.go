package lib

import "net/http"

type CustomError struct {
	Message  string
	Field    string
	Code     int
	HTTPCode int
}

func (err CustomError) Error() string {
	return err.Message
}

var (
	ErrorForbidden = CustomError{
		Message:  "Forbidden",
		Code:     1000,
		HTTPCode: http.StatusForbidden,
	}

	ErrorInvalidParameter = CustomError{
		Message:  "Invalid Parameter",
		Code:     1001,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorWrongUsrnamePassword = CustomError{
		Message:  "Wrong Username or Password",
		Code:     1002,
		HTTPCode: http.StatusUnprocessableEntity,
	}
)
