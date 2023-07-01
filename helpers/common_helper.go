package helpers

import (
	"regexp"
	"time"
)

func ValidateMobileNumber(number string) bool {
	// Define a regular expression pattern for a mobile number
	// Modify the pattern according to the format of mobile numbers in your specific region
	pattern := `^\+[1-9]\d{1,14}$`

	// Create a regular expression object
	regex := regexp.MustCompile(pattern)

	// Match the number against the regular expression
	return regex.MatchString(number)
}

func ValidateDateOfBirth(dateString string) bool {
	// Define the expected date format
	layout := "2006-01-02" // yyyy-mm-dd

	// Parse the date string
	date, err := time.Parse(layout, dateString)
	if err != nil {
		return false // Parsing error, not a valid date
	}
	// Check if the parsed date is valid
	now := time.Now()

	return date.Before(now)
}
