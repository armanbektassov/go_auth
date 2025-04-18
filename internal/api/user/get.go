package user

import (
	"context"
	"log"

	"github.com/armanbektassov/go_auth/internal/converter"
	desc "github.com/armanbektassov/go_auth/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, title: %s, body: %s, created_at: %v, updated_at: %v\n", userObj.ID, userObj.Info.Name, userObj.Info.Email, userObj.CreatedAt, userObj.UpdatedAt)

	return &desc.GetResponse{
		User: converter.ToUserFromService(userObj),
	}, nil
}
