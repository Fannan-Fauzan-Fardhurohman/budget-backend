package common

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ApiResponse map[string]any
type ValidationError struct {
	Error     string `json:"error"`
	Key       string `json:"key"`
	Condition string `json:"condition"`
}

type JSONSuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONFailedValidationResponse struct {
	Success bool               `json:"success"`
	Errors  []*ValidationError `json:"error"`
	Message string             `json:"message"`
}

type JSONErrorResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Errors  ValidationError `json:"error"`
}

func SendSuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, JSONSuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendFailedValidationResponse(c echo.Context, errors []*ValidationError) error {
	return c.JSON(http.StatusBadRequest, JSONFailedValidationResponse{
		Success: false,
		Errors:  errors,
		Message: "Validation Failed",
	})
}

func SendErrorResponse(c echo.Context, message string, statusCode int) error {
	return c.JSON(statusCode, JSONErrorResponse{
		Success: false,
		Message: message,
	})
}

func SendBadRequestResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, message, http.StatusBadRequest)
}
func SendNotFoundResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, message, http.StatusNotFound)
}

func SendInternalServerErrorResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, message, http.StatusInternalServerError)
}

func SendUnauthorizationErrorResponse(c echo.Context, message string) error {

	return SendErrorResponse(c, message, http.StatusUnauthorized)
}
