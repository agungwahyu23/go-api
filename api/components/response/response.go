package response

import "net/http"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(message string, data interface{}) (int, APIResponse) {
	return http.StatusOK, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func Created(message string, data interface{}) (int, APIResponse) {
	return http.StatusCreated, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func Error(status int, message string, err interface{}) (int, APIResponse) {
	return status, APIResponse{
		Success: false,
		Message: message,
		Errors:  err,
	}
}

func ValidationError(err interface{}) (int, APIResponse) {
	return http.StatusUnprocessableEntity, APIResponse{
		Success: false,
		Message: "Validation error",
		Errors:  err,
	}
}
