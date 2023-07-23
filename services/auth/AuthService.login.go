package auth

import (
	"gotest/common/errors"
	"gotest/models"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User   *models.User `json:"user"`
	Tokens *Tokens      `json:"tokens"`
}

func (service *AuthService) Login(request *LoginRequest) (*LoginResponse, error) {
	userDto := service.userRepository.Find(request.Username)
	if userDto == nil {
		return nil, &errors.APIError{Message: "User not found"}
	}

	if (request.Username == userDto.Username) && service.CheckPasswordHash(request.Password, userDto.PasswordHash) {
		accessToken, err := service.generateJwtToken(userDto)

		if err != nil {
			return nil, err
		}

		var userModel = models.User{Id: userDto.Id, Username: userDto.Username, Email: userDto.Email, Photo: userDto.Photo}

		return &LoginResponse{User: &userModel, Tokens: &Tokens{AccessToken: accessToken, RefreshToken: "refresh-token"}}, nil
	} else {
		return nil, &errors.APIError{Message: "Username or password not corect"}
	}
}
