package auth

import (
	"context"
	"github.com/armanbektassov/go_auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"github.com/armanbektassov/go_auth/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(s.tokenConfig.RefreshToken()))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	accessToken, err := utils.GenerateToken(model.UserInfo{
		Name: claims.Username,
		Role: claims.Role,
	},
		[]byte(s.tokenConfig.RefreshToken()),
		time.Duration(s.tokenConfig.AccessTokenExpiration())*time.Minute,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return accessToken, nil
}
