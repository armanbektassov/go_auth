package user

import (
	"github.com/armanbektassov/go_auth/internal/service"
	desc "github.com/armanbektassov/go_auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
