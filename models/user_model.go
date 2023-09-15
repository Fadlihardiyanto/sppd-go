package models

type CreateUserRequest struct {
	IDUser       string `json:"user_id" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Email        string `json:"email" validate:"email"`
	Nama         string `json:"nama" validate:"required"`
	NomorTelepon string `json:"nomor_telepon" validate:"required,numeric"`
	Alamat       string `json:"alamat" validate:"required"`
}

type CreateUserResponse struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Nama         string `json:"nama"`
	NomorTelepon string `json:"nomor_telepon"`
	Alamat       string `json:"alamat"`
}
