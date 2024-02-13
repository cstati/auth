package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/hse-experiments-platform/auth/internal/pkg/storage/db"
	errs "github.com/hse-experiments-platform/auth/pkg/utils/web/errors"
	"github.com/jackc/pgx/v5"
)

type LoginWithGoogleRequest struct {
	GoogleOAuthToken string `json:"google_oauth_token"`
}

type LoginWithGoogleResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

// LoginWithGoogle godoc
// @Summary      Try login with Google OAuth2 token
// @Description  Get userID and paseto token by google oauth2 token
// @Accept       json
// @Produce      json
// @Param        body  body    LoginWithGoogleRequest  true  "body"
// @Success      200  {object}  LoginWithGoogleResponse
// @Failure      401  {object}  errors.CodedError
// @Failure      500  {object}  errors.CodedError
// @Router       /login/google [post]
func (s *AuthService) LoginWithGoogle(ctx context.Context, headers http.Header, r *http.Request) (*LoginWithGoogleResponse, error) {
	var request LoginWithGoogleRequest

	if r.ContentLength == 0 {
		return nil, errs.NewCodedError(http.StatusBadRequest, fmt.Errorf("no token provided"))
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, errs.InternalError(err)
	} else if request.GoogleOAuthToken == "" {
		return nil, errs.NewCodedError(http.StatusBadRequest, fmt.Errorf("no token provided"))
	}

	slog.Debug("got google token", slog.String("token", request.GoogleOAuthToken))

	googleInfo, err := s.google.GetIDAndEmail(ctx, request.GoogleOAuthToken)
	if err != nil {
		return nil, fmt.Errorf("s.google.GetIDAndEmail(%s): %w", request.GoogleOAuthToken, err)
	}

	slog.Debug("got google info", slog.Any("info", googleInfo))

	user, err := s.db.GetUserByExternalID(ctx, googleInfo.Id)
	if errors.Is(err, pgx.ErrNoRows) {
		if user.ID, err = s.db.CreateUser(ctx, db.CreateUserParams{
			GoogleID: googleInfo.Id,
			Email:    googleInfo.Email,
		}); err != nil {
			return nil, fmt.Errorf("s.db.CreateUser(id=%s, email=%s): %w", googleInfo.Id, googleInfo.Email, err)
		}
	} else if err != nil {

		return nil, fmt.Errorf("s.db.GetUserIDByExternalID(%s): %w", googleInfo.Id, err)
	}
	token, err := s.tokenProvider.CreateToken(user.ID, nil, time.Hour*24)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}

	return &LoginWithGoogleResponse{UserID: user.ID, Token: token}, nil
}
