package helpers

import "testing"

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
	}{
		// valid URLs
		{"Valid HTTPS", "https://example.com", true},
		{"Valid HTTP", "http://example.com", true},
		{"Subdomain", "https://sub.example.co.uk", true},
		{"Path and query", "https://example.com/search?q=golang", true},

		// invalid URLs
		{"Missing scheme", "example.com", false},
		{"Unsupported scheme", "ftp://example.com", false},
		{"Just text", "hello world", false},
		{"Missing host", "http://", false},
		{"No scheme, has path", "/some/path", false},
		{"IP without scheme", "127.0.0.1", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsValidURL(test.input)
			if got != test.valid {
				t.Errorf("IsValidURL(%q) = %v; want %v", test.input, got, test.valid)
			}
		})
	}
}
