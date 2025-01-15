package middleware

import (
	"encoding/json"
	"net/http"
)

// ResponseFormat defines a standard format for API responses
type ResponseFormat struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Response handles both success and error responses in a unified way
func Response(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set the response status code
	w.WriteHeader(statusCode)

	// Determine status based on status code
	var status string
	if statusCode >= 200 && statusCode < 300 {
		status = "success" // 2xx success codes
	} else {
		status = "error" // 4xx and 5xx error codes
	}

	// Create the response body
	response := ResponseFormat{
		Status:  status,
		Message: message,
		Data:    data,
	}

	// Encode the response into JSON and send it to the client
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
