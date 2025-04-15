package utils

import (
	"net/url"
	"regexp"
)

// format and scheme/host check
func IsValidURL(raw string) bool {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}

	// Ensure scheme is http/https and host exists
	re := regexp.MustCompile(`^https?://`)
	return re.MatchString(raw) && parsed.Host != ""
}
