package model

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
