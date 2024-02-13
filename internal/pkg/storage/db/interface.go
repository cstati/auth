package db

import "context"

type Storage interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	GetUserByExternalID(ctx context.Context, googleID string) (User, error)
}
