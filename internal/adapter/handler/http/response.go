package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	port "github.com/GeekGawd/logtracer/internal/core/port"
)

var errorStatusMap = map[error]int{
	port.ErrDataNotFound:               http.StatusNotFound,
	port.ErrConflictingData:            http.StatusConflict,
	port.ErrInvalidCredentials:         http.StatusUnauthorized,
	port.ErrUnauthorized:               http.StatusUnauthorized,
	port.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	port.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	port.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	port.ErrInvalidToken:               http.StatusUnauthorized,
	port.ErrExpiredToken:               http.StatusUnauthorized,
	port.ErrForbidden:                  http.StatusForbidden,
	port.ErrNoUpdatedData:              http.StatusBadRequest,
	port.ErrInsufficientStock:          http.StatusBadRequest,
	port.ErrInsufficientPayment:        http.StatusBadRequest,
}

// response represents a response body format
type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// newResponse is a helper function to create a response body
func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// errorResponse represents an error response body format
type errorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error message"`
}

// newErrorResponse is a helper function to create an error response body
func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
	}
}

// validationError sends an error response for some specific request validation error
func validationError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, err)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errRsp := newErrorResponse(err.Error())

	ctx.JSON(statusCode, errRsp)
}

// handleAbort sends an error response and aborts the request with the specified status code and error message
func handleAbort(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	rsp := newErrorResponse(err.Error())
	ctx.AbortWithStatusJSON(statusCode, rsp)
}

// handleSuccess sends a success response with the specified status code and optional data
func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}