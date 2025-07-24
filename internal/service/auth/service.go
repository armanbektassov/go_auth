package auth

import (
	"github.com/armanbektassov/go_auth/internal/config"
	"github.com/armanbektassov/go_auth/internal/repository"
	"github.com/armanbektassov/go_auth/internal/service"
	"github.com/armanbektassov/platform_common/pkg/client/db"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
	tokenConfig    config.TokenConfig
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
	tokenConfig config.TokenConfig,
) service.AuthService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
		tokenConfig:    tokenConfig,
	}
}
