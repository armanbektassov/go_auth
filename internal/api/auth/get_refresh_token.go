package auth

import (
	"context"
	"log"

	desc "github.com/armanbektassov/go_auth/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.RefreshTokenResponse, error) {
	token, err := i.authService.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, err
	}

	log.Printf("successfully refresh token: %s", token)

	return &desc.RefreshTokenResponse{
		RefreshToken: token,
	}, nil
}
