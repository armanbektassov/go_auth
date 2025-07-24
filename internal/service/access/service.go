package access

import (
	"github.com/armanbektassov/go_auth/internal/config"
	"github.com/armanbektassov/go_auth/internal/service"
)

type serv struct {
	tokenConfig config.TokenConfig
}

func NewService(
	tokenConfig config.TokenConfig,
) service.AccessService {
	return &serv{
		tokenConfig: tokenConfig,
	}
}
