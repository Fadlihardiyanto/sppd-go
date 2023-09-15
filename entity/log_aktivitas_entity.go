package entity

type LogAktivitasEntity struct {
	IDLog            string `gorm:"primaryKey;not null" json:"log_id"`
	UserID           string `gorm:"type:varchar(255);not null" json:"user_id"`
	TanggalAktivitas string `gorm:"type:datetime;not null" json:"tanggal_aktivitas"`
	JenisAktivitas   string `gorm:"type:varchar(255);not null" json:"jenis_aktivitas"`
}

func (LogAktivitasEntity) TableName() string {
	return "log_aktivitas"
}
