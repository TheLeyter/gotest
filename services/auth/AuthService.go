package auth

import "gotest/repositories"

type AuthService struct {
	userRepository *repositories.UserRepository
}

type Tokens struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

func NewAuthService(repository *repositories.UserRepository) *AuthService {
	authService := AuthService{userRepository: repository}
	return &authService
}
