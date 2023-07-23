package auth

import "gotest/models"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User   *models.User `json:"user"`
	Tokens *Tokens      `json:"tokens"`
}

func (service *AuthService) Login(request *LoginRequest) *LoginResponse {
	user := service.userRepository.Find(request.Username)
	if user == nil {
		return nil
	}

	return &LoginResponse{User: user, Tokens: &Tokens{AccessToken: "access-token", RefreshToken: "refresh-token"}}
}
