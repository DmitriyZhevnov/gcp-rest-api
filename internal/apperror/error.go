package apperror

import (
	"encoding/json"
)

var (
	errNotFound         = newAppError("Not found")
	internalServerError = newAppError("Internal Server error")
	badRequestError     = newAppError("Bad Request")
	unprocessableEntity = newAppError("Unprocessable Entity")
)

type appError struct {
	Err              error  `json:"-"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
	Code             string `json:"code"`
}

func newAppError(message string) *appError {
	return &appError{
		Message: message,
	}
}

func (err *appError) fillFields(developerMessage, code string) *appError {
	err.DeveloperMessage = developerMessage
	err.Code = code
	return err
}

func NewErrNotFound(developerMessage, code string) *appError {
	err := errNotFound
	return err.fillFields(developerMessage, code)
}

func NewInternalServerError(developerMessage, code string) *appError {
	err := internalServerError
	return err.fillFields(developerMessage, code)
}

func NewBadRequestError(developerMessage, code string) *appError {
	err := badRequestError
	return err.fillFields(developerMessage, code)
}

func NewUnprocessableEntityError(developerMessage, code string) *appError {
	err := unprocessableEntity
	return err.fillFields(developerMessage, code)
}

func (a *appError) Error() string {
	return a.Message
}

func (a *appError) Unwrap() error {
	return a.Err
}

func (a *appError) Marshal() []byte {
	marshal, err := json.Marshal(a)
	if err != nil {
		return nil
	}
	return marshal
}
