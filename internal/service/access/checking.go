package access

import (
	"context"
	"errors"
	"github.com/armanbektassov/go_auth/internal/utils"
	"google.golang.org/grpc/metadata"
	"strings"
)

func (s *serv) Checking(ctx context.Context, endpoint string) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return false, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], s.tokenConfig.AuthPrefix()) {
		return false, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], s.tokenConfig.AuthPrefix())

	claims, err := utils.VerifyToken(accessToken, []byte(s.tokenConfig.AccessToken()))
	if err != nil {
		return false, errors.New("access token is invalid")
	}

	accessibleMap, err := utils.AccessibleRoles()
	if err != nil {
		return false, errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[endpoint]
	if !ok {
		return true, nil
	}

	if role == utils.GetRoleByUserRole(claims.Role) {
		return true, nil
	}

	return false, errors.New("access denied")
}
