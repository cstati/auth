package auth

import (
	"github.com/hse-experiments-platform/auth/internal/pkg/storage/db"
	"github.com/hse-experiments-platform/auth/internal/pkg/storage/google"
	"github.com/hse-experiments-platform/library/pkg/utils/token"
)

// typecheck
var _ Service = &AuthService{}

type AuthService struct {
	google        google.Storage
	db            db.Storage
	tokenProvider token.Maker
}

func NewService(google google.Storage, db db.Storage, tokenProvider token.Maker) *AuthService {
	return &AuthService{
		google:        google,
		db:            db,
		tokenProvider: tokenProvider,
	}
}
