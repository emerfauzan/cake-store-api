package model

type User struct {
	Id                uint   `json:"id"`
	Name              string `json:"name"`
	Username          string `json:"username"`
	EncryptedPassword string `json:"encrypted_password"`
}
