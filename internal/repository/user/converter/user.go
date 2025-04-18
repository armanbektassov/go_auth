package converter

import (
	"github.com/armanbektassov/go_auth/internal/model"
	modelRepo "github.com/armanbektassov/go_auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(info modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Role:     info.Role,
		Password: info.Password,
	}
}
