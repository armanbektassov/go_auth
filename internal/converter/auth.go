package converter

import (
	"github.com/armanbektassov/go_auth/internal/model"
	desc "github.com/armanbektassov/go_auth/pkg/auth_v1"
)

func ToAuthFromService(info model.Auth) *desc.LoginInfo {
	return &desc.LoginInfo{
		Username: info.Username,
		Password: info.Password,
	}
}

func ToAuthFromDesc(info *desc.LoginInfo) *model.Auth {
	return &model.Auth{
		Username: info.Username,
		Password: info.Password,
	}
}
