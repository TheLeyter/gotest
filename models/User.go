package models

type User struct {
	Id           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Salt         string `json:"salt"`
	PasswordHash string `json:"password"`
}
