package auth

import (
	"gotest/common/envs"
	"gotest/repositories"
)

type AuthService struct {
	userRepository *repositories.UserRepository
	config         *envs.ApiConfig
}

func NewAuthService(repository *repositories.UserRepository, config *envs.ApiConfig) *AuthService {
	authService := AuthService{userRepository: repository, config: config}
	return &authService
}
