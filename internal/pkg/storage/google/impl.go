package google

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	errs "github.com/hse-experiments-platform/library/pkg/utils/web/errors"
	"google.golang.org/api/googleapi"
	oauth2api "google.golang.org/api/oauth2/v2"
)

const (
	googleURL = "https://www.googleapis.com/userinfo/v2/me"
)

type googleImpl struct {
}

func NewStorage() Storage {
	return &googleImpl{}
}

func (google *googleImpl) GetIDAndEmail(ctx context.Context, googleToken string) (*oauth2api.Userinfo, error) {
	service, err := oauth2api.New(&http.Client{})
	if err != nil {
		return nil, fmt.Errorf("oauth2.NewService: %w", err)
	}

	req := service.Userinfo.V2.Me.Get()
	req.Header().Add("Authorization", "Bearer "+googleToken)
	info, err := req.Do()
	if err != nil {
		var googleErr *googleapi.Error
		if errors.As(err, &googleErr) {
			return nil, errs.NewCodedError(googleErr.Code, err)
		} else {
			return nil, err
		}
	}

	return info, nil
}
