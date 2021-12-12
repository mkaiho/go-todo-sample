package usecase

import "fmt"

type errorCode int

const (
	ErrCodeNotFoundEntity errorCode = iota
	ErrCodeInvalidUsecaseInput
	ErrCodeUnkown
)

type UseCaseError interface {
	Code() errorCode
	Error() string
}

func NewErrNotFoundEntity(name string, param interface{}) UseCaseError {
	return &useCaseError{
		code:    ErrCodeNotFoundEntity,
		message: fmt.Sprintf("%s entity is not found. param: %s", name, param),
	}
}

func NewErrInvalidUsecaseInput(name string, fields map[string]interface{}) UseCaseError {
	return &useCaseError{
		code:    ErrCodeInvalidUsecaseInput,
		message: fmt.Sprintf("%s input is invalid: %v", name, fields),
	}
}

func NewErrUnknown(err error) UseCaseError {
	return &useCaseError{
		code:    ErrCodeUnkown,
		message: fmt.Sprintf("unknown error is occured: %s", err),
	}
}

type useCaseError struct {
	code    errorCode
	message string
}

func (e *useCaseError) Code() errorCode {
	return e.code
}

func (e *useCaseError) Error() string {
	return e.message
}

func (e *useCaseError) IsCodeEqualTo(code errorCode) bool {
	return e.code == code
}
