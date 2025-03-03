package utils

import "regexp"

func ValidatePassword(password string) bool {
	// validate the password
	if len(password) < 8 {
		return false
	}

	// Define regex patterns for each requirement
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)

	// Return true only if all conditions are met
	return hasLower && hasUpper && hasDigit && hasSpecial
}