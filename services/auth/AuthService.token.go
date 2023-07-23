package auth

import (
	"encoding/json"
	"fmt"
	"gotest/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

func (service *AuthService) generateJwtToken(user *models.UserDto) (string, error) {
	secret := []byte(service.config.Secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   json.Number(strconv.FormatInt(time.Now().Add(time.Minute*time.Duration(30)).Unix(), 10)),
		"iat":   json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		"id":    json.Number(strconv.FormatInt(int64(user.Id), 10)),
		"name":  user.Username,
		"email": user.Email,
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
