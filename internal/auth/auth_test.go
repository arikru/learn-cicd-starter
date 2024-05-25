package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test when the Authorization header is missing
	t.Run("No Authorization Header", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
	})

	// Test when the Authorization header is malformed
	t.Run("Malformed Authorization Header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer some_api_key")
		_, err := GetAPIKey(headers)
		if err == nil || err.Error() != "malformed authorization header" {
			t.Fatalf("expected error 'malformed authorization header', got %v", err)
		}
	})

	// Test when the Authorization header is correctly formatted
	t.Run("Valid Authorization Header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey some_api_key")
		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if apiKey != "some_api_key" {
			t.Fatalf("expected apiKey 'some_api_key', got %v", apiKey)
		}
	})
}
