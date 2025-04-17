package helpers

import (
	"net/url"
	"regexp"
)

// checks if the given string is a valid URL
func IsValidURL(raw string) bool {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}

	// Ensure scheme is http/https and host exists
	re := regexp.MustCompile(`^https?://`)
	return re.MatchString(raw) && parsed.Host != ""
}
