package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract api key from header
// Example: Authorization: ApiKey {apiKey}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("sem informação de autenticação")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("header de autenticação com com mal formação")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("argumento inexistente no header de autenticação")
	}

	return vals[1], nil
}
