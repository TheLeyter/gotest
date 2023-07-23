package auth

import (
	"gotest/common/errors"
	"gotest/models"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	User   *models.User `json:"user"`
	Tokens *Tokens      `json:"tokens"`
}

func (service *AuthService) Register(request *RegisterRequest) (*RegisterResponse, error) {
	// findUser := service.userRepository.Find(request.Username)

	// fmt.Println(findUser)

	// if findUser != nil {
	// 	fmt.Println("-------------------")
	// 	return nil
	// }

	hashPassword, err := service.HashPassword(request.Password)

	if err != nil {
		return nil, &errors.APIError{Message: "Incorect password"}
	}

	userDto, err := service.userRepository.Create(request.Username, request.Email, hashPassword)

	if err != nil {
		return nil, err
	}

	var userModel = models.User{Id: userDto.Id, Username: userDto.Username, Email: userDto.Email, Photo: userDto.Photo}

	accessToken, err := service.generateJwtToken(userDto)

	if err != nil {
		return nil, &errors.APIError{Message: "Token can't created"}
	}

	var tokens = Tokens{AccessToken: accessToken, RefreshToken: "refresh-token"}

	return &RegisterResponse{User: &userModel, Tokens: &tokens}, nil
}
