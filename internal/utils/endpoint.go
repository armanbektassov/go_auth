package utils

import "github.com/armanbektassov/go_auth/internal/model"

var accessibleRoles map[string]string

func AccessibleRoles() (map[string]string, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]string)

		accessibleRoles[model.ExamplePath] = "Admin"
	}

	return accessibleRoles, nil
}

func GetRoleByUserRole(role int64) string {
	switch role {
	case 0:
		return "User"
	case 1:
		return "Admin"
	default:
		return "Unknown"
	}
}
