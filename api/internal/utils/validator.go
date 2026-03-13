package utils

import (
	"regexp"
	"strings"
)

// IsValidEmail validates an email address format
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsEmpty checks if a string is empty or contains only whitespace
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// ValidateRequired checks if required fields are present
func ValidateRequired(fields map[string]string) []string {
	var errors []string

	for field, value := range fields {
		if IsEmpty(value) {
			errors = append(errors, field+" is required")
		}
	}

	return errors
}

// ValidateCreateUser validates user creation input
func ValidateCreateUser(name, email string) []string {
	var errors []string

	if IsEmpty(name) {
		errors = append(errors, "name is required")
	}

	if IsEmpty(email) {
		errors = append(errors, "email is required")
	} else if !IsValidEmail(email) {
		errors = append(errors, "email format is invalid")
	}

	return errors
}
