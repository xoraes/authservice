package models

import (
	"fmt"
	"net/http"
)

const UNAUTHORIZED = http.StatusUnauthorized
const BADREQUEST = http.StatusBadRequest

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *AppError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.Code, r.Message)
}

func appErr(message string, code int) *AppError {
	return &AppError{Message: message, Code: code}
}
