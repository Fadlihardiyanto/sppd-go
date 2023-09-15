package entity

import "time"

type PersetujuanEntity struct {
	IDPersetujuan string `gorm:"primaryKey;not null" json:"persetujuan_id"`
	SuratID       string `gorm:"type:varchar(255);not null" json:"surat_id"`
	// SuratPerjalananDinasEntity SuratPerjalananDinasEntity `gorm:"references:IDSurat"`
	UserID string `gorm:"type:varchar(255);not null" json:"user_id"`
	// User               UserEntity `gorm:"foreignKey:UserID" json:"user_entity"`
	TanggalPersetujuan time.Time `gorm:"type:datetime;not null" json:"tanggal_persetujuan"`
	StatusPersetujuan  int       `gorm:"type:int;not null" json:"status_persetujuan"`
}

func (PersetujuanEntity) TableName() string {
	return "persetujuan"
}
