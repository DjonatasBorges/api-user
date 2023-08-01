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
	ErrUserNotFound = NewAppError(1001, "%s not found")
	ErrInvalidUUID  = &AppError{Code: 1002, Message: "Invalid UUID"}
)
