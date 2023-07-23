package models

type UserDto struct {
	Id           uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Username     string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string
	Photo        string
}

func (UserDto) TableName() string {
	return "users"
}
