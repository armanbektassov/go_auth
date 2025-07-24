package auth

import (
	"context"
	"github.com/armanbektassov/go_auth/internal/utils"
	"time"

	"github.com/armanbektassov/go_auth/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) Login(ctx context.Context, info *model.Auth) (string, error) {
	user, err := s.userRepository.GetByUsername(ctx, info.Username)
	if err != nil {
		return "", err
	}

	if user.Info.Password != info.Password {
		return "", errors.New("Incorrect username or password")
	}

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Name: user.Info.Name,
		Role: user.Info.Role,
	},
		[]byte(s.tokenConfig.RefreshToken()),
		time.Duration(s.tokenConfig.RefreshTokenExpiration())*time.Minute,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
