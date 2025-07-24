package auth

import (
	"context"
	"log"

	"github.com/armanbektassov/go_auth/internal/converter"
	desc "github.com/armanbektassov/go_auth/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.RefreshTokenResponse, error) {
	token, err := i.authService.Login(ctx, converter.ToAuthFromDesc(req.GetLoginInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("successfully login with token: %s", token)

	return &desc.RefreshTokenResponse{
		RefreshToken: token,
	}, nil
}
