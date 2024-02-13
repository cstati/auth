package google

import (
	"context"

	oauth2api "google.golang.org/api/oauth2/v2"
)

type Storage interface {
	GetIDAndEmail(ctx context.Context, googleToken string) (*oauth2api.Userinfo, error)
}
