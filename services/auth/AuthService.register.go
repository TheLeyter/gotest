package auth

import "gotest/models"

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	User   *models.User `json:"user"`
	Tokens *Tokens      `json:"tokens"`
}

func (service *AuthService) Register(request *RegisterRequest) *RegisterResponse {
	user := service.userRepository.Find(request.Username)
	if user != nil {
		return nil
	}

	return &RegisterResponse{User: user, Tokens: &Tokens{AccessToken: "access-token", RefreshToken: "refresh-token"}}
}
