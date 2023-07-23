package repositories

import (
	"gotest/common/db"
	"gotest/models"
)

type UserRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) Create(username, email, passwordHash string) (*models.UserDto, error) {
	user := models.UserDto{Username: username, Email: email, PasswordHash: passwordHash}
	res := repository.db.DB.Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (repository *UserRepository) Find(name string) *models.UserDto {
	var user models.UserDto
	repository.db.DB.Where("username = ?", name).Or("email = ?", name).First(&user)
	return &user
}
