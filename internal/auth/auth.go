package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the headers of an HTTP request.
// Example:
// Authorization: ApiKey {api_key}
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no Authorization header found")
	}

	values := strings.Split(authHeader, " ")
	if len(values) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	if values[0] != "ApiKey" {
		return "", errors.New("invalid Authorization header")
	}

	return values[1], nil
}
