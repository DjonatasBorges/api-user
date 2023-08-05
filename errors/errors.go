package errors

import (
	"fmt"
	"log"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	msg := fmt.Sprintf("Error code %d: %s", e.Code, e.Message)
	log.Println(msg)
	return msg
}

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (e *AppError) WithArgs(args ...interface{}) *AppError {
	message := fmt.Sprintf(e.Message, args...)
	return &AppError{Code: e.Code, Message: message}
}

var (
	ErrUserNotFound          = NewAppError(1001, "%s not found")
	ErrFieldCannotBeNull     = NewAppError(1002, "%s cannot be null, empty or with an invalid value.")
	ErrInvalidUUID           = NewAppError(1003, "Invalid UUID")
	ErrMissingAuthToken      = NewAppError(1004, "Missing authentication token")
	ErrInvalidAuthHeader     = NewAppError(1005, "Invalid authorization header")
	ErrInvalidSigningMethod  = NewAppError(1006, "Invalid signing method")
	ErrInvalidCredentials    = NewAppError(1007, "Invalid credentials")
	ErrUserNotRegistered     = NewAppError(1008, "User not registered")
	ErrInvalidTokenSignature = NewAppError(1009, "Invalid token signature")
)
