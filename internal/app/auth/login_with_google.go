package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/cstati/auth/internal/pkg/storage/db"
	pb "github.com/cstati/auth/pkg/auth"
	"github.com/jackc/pgx/v5"
)

type LoginWithGoogleRequest struct {
	GoogleOAuthToken string `json:"google_oauth_token"`
}

type LoginWithGoogleResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func (s *Service) GoogleLogin(ctx context.Context, r *pb.GoogleLoginRequest) (*pb.GoogleLoginResponse, error) {
	slog.Debug("got google token", slog.String("token", r.GoogleOauthToken))

	googleInfo, err := s.google.GetIDAndEmail(ctx, r.GoogleOauthToken)
	if err != nil {
		return nil, fmt.Errorf("s.google.GetIDAndEmail(%s): %w", r.GoogleOauthToken, err)
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
	token, err := s.tokenProvider.CreateToken(user.ID, nil, time.Hour*24*7)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}

	roles, err := s.db.GetUserRolesByID(ctx, user.ID)

	return &pb.GoogleLoginResponse{
		UserId: user.ID,
		Token:  token,
		Roles:  roles,
	}, nil
}
