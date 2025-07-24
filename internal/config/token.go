package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	refreshTokenEnvName           = "REFRESH_TOKEN_SECRET_KEY"
	accessTokenEnvName            = "ACCESS_TOKEN_SECRET_KEY"
	authPrefixEnvName             = "AUTH_PREFIX"
	refreshTokenExpirationEnvName = "REFRESH_TOKEN_EXPIRATION"
	accessTokenExpirationEnvName  = "ACCESS_TOKEN_EXPIRATION"
)

type TokenConfig interface {
	RefreshToken() string
	AccessToken() string
	AuthPrefix() string
	RefreshTokenExpiration() int
	AccessTokenExpiration() int
}

type tokenConfig struct {
	refreshToken           string
	accessToken            string
	authPrefix             string
	refreshTokenExpiration int
	accessTokenExpiration  int
}

func NewTokenConfig() (TokenConfig, error) {
	refreshToken := os.Getenv(refreshTokenEnvName)
	if len(refreshToken) == 0 {
		return nil, errors.New("refreshToken not found")
	}

	accessToken := os.Getenv(accessTokenEnvName)
	if len(accessToken) == 0 {
		return nil, errors.New("accessToken not found")
	}

	authPrefix := os.Getenv(authPrefixEnvName)
	if len(authPrefix) == 0 {
		return nil, errors.New("authPrefix not found")
	}

	refreshTokenExpiration := os.Getenv(refreshTokenExpirationEnvName)
	if len(refreshTokenExpiration) == 0 {
		return nil, errors.New("refreshTokenExpiration not found")
	}

	rt, err := strconv.Atoi(refreshTokenExpiration)
	if err != nil {
		return nil, errors.New("error converting refreshTokenExpiration to int")
	}

	accessTokenExpiration := os.Getenv(accessTokenExpirationEnvName)
	if len(accessTokenExpiration) == 0 {
		return nil, errors.New("accessTokenExpiration not found")
	}

	at, err := strconv.Atoi(accessTokenExpiration)
	if err != nil {
		return nil, errors.New("error converting refreshTokenExpiration to int")
	}

	return &tokenConfig{
		refreshToken:           refreshToken,
		accessToken:            accessToken,
		authPrefix:             authPrefix,
		refreshTokenExpiration: rt,
		accessTokenExpiration:  at,
	}, nil
}

func (cfg *tokenConfig) RefreshToken() string {
	return cfg.refreshToken
}

func (cfg *tokenConfig) AccessToken() string {
	return cfg.accessToken
}

func (cfg *tokenConfig) AuthPrefix() string {
	return cfg.authPrefix
}

func (cfg *tokenConfig) RefreshTokenExpiration() int {
	return cfg.refreshTokenExpiration
}

func (cfg *tokenConfig) AccessTokenExpiration() int {
	return cfg.accessTokenExpiration
}
