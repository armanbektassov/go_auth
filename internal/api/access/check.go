package access

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	desc "github.com/armanbektassov/go_auth/pkg/access_v1"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	_, err := i.accessService.Checking(ctx, req.EndpointAddress)
	if err != nil {
		return nil, err
	}

	log.Print("access grant")

	return &emptypb.Empty{}, nil
}
