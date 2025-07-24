package service

import (
	"context"

	"github.com/armanbektassov/go_auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}

type AuthService interface {
	Login(ctx context.Context, info *model.Auth) (string, error)
	GetRefreshToken(ctx context.Context, oldToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}

type AccessService interface {
	Checking(ctx context.Context, endpoint string) (bool, error)
}
