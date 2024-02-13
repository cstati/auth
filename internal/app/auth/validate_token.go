package auth

import (
	"context"
	"net/http"
	"strings"

	errs "github.com/hse-experiments-platform/auth/pkg/utils/web/errors"
)

type UserInfo struct {
	UserID int64    `json:"user_id"`
	Roles  []string `json:"roles"`
}

// CheckToken godoc
// @Summary      Validate user's token
// @Description  Validate user's token and if correct return UserInfo
// @Produce      json
// @Param        Authorization  header    string  false  "Paseto encrypted token" example(Bearer v2.local.ABCDEFG)
// @Success      200  {object}  UserInfo
// @Failure      403  {object}  errors.CodedError
// @Failure      500  {object}  errors.CodedError
// @Router       /validate [get]
func (s *AuthService) ValidateToken(_ context.Context, _ http.Header, r *http.Request) (*UserInfo, error) {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	payload, err := s.tokenProvider.VerifyToken(token)
	if err != nil {
		return nil, errs.NewCodedError(http.StatusForbidden, err)
	}

	return &UserInfo{
		UserID: payload.UserID,
		Roles:  payload.Roles,
	}, nil
}
