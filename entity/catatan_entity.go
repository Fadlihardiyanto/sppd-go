package entity

import "time"

type CatatanEntity struct {
	IDCatatan string `gorm:"primaryKey;not null" json:"catatan_id"`
	SuratID   string `gorm:"type:varchar(255);not null" json:"surat_id"`
	// SuratPerjalananDinasEntity SuratPerjalananDinasEntity `gorm:"references:IDSurat"`
	TanggalCatatan time.Time `gorm:"type:datetime;not null" json:"tanggal_catatan"`
	IsiCatatan     string    `gorm:"type:varchar(255);not null" json:"isi_catatan"`
}

func (CatatanEntity) TableName() string {
	return "catatan"
}
