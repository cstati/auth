package auth

import (
	"context"
	"net/http"
)

type Service interface {
	LoginWithGoogle(context.Context, http.Header, *http.Request) (*LoginWithGoogleResponse, error)
	Logout(context.Context, http.Header, *http.Request) (string, error)
	ValidateToken(context.Context, http.Header, *http.Request) (*UserInfo, error)
}
