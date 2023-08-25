package security

import (
	"errors"
	"net/http"
	"os"
)

type AuthorizationResponse struct {
	Valid   bool
	Message error
}

func AuthorizeRequest(r *http.Request) *AuthorizationResponse {
	authorizationToken := os.Getenv("AUTHORIZATION_TOKEN")
	if authorizationTokenInRequest := r.Header.Get("Authorization"); authorizationToken == authorizationTokenInRequest {
		return &AuthorizationResponse{
			Valid:   true,
			Message: nil,
		}
	}

	return &AuthorizationResponse{
		Valid:   false,
		Message: errors.New("unauthorized token"),
	}
}
