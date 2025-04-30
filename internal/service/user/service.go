package user

import (
	"github.com/armanbektassov/go_auth/internal/repository"
	"github.com/armanbektassov/go_auth/internal/service"
	"github.com/armanbektassov/platform_common/pkg/client/db"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
