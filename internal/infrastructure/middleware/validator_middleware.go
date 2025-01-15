package middleware

import (
	"github.com/go-playground/validator/v10"
)

// Global validator instance
var Validate *validator.Validate

func init() {
	// Initialize the validator instance
	Validate = validator.New()
}
