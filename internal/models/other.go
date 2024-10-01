package models

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

type Response struct {
	Message interface{} `json:"message"`
}
