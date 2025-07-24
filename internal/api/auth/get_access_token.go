package auth

import (
	"context"
	"log"

	desc "github.com/armanbektassov/go_auth/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.AccessTokenResponse, error) {
	token, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	log.Printf("successfully create access token: %s", token)

	return &desc.AccessTokenResponse{
		AccessToken: token,
	}, nil
}
