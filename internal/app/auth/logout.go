package auth

import (
	"context"
	"net/http"
)

func (s *AuthService) Logout(ctx context.Context, headers http.Header, r *http.Request) (string, error) {
	panic("not implemented yet")
}
