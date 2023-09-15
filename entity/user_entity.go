package entity

type UserEntity struct {
	IDUser       string `gorm:"primaryKey;not null" json:"user_id"`
	Username     string `gorm:"type:varchar(255);not null" json:"username"`
	Password     string `gorm:"type:varchar(255);not null" json:"password"`
	Email        string `gorm:"type:varchar(255);not null" json:"email"`
	Nama         string `gorm:"type:varchar(255);not null" json:"nama"`
	NomorTelepon string `gorm:"type:varchar(255);not null" json:"nomor_telepon"`
	Alamat       string `gorm:"type:varchar(255);not null" json:"alamat"`
}

func (UserEntity) TableName() string {
	return "users"
}
