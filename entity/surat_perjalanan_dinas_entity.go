package entity

type SuratPerjalananDinasEntity struct {
	IDSurat          string `gorm:"primaryKey;not null" json:"surat_id"`
	NomorSurat       string `gorm:"type:varchar(255);not null" json:"nomor_surat"`
	TanggalSurat     string `gorm:"type:varchar(255);not null" json:"tanggal_surat"`
	TanggalBerangkat string `gorm:"type:varchar(255);not null" json:"tanggal_berangkat"`
	TanggalKembali   string `gorm:"type:varchar(255);not null" json:"tanggal_kembali"`
	Tujuan           string `gorm:"type:varchar(255);not null" json:"tujuan"`
	Anggaran         string `gorm:"type:varchar(255);not null" json:"anggaran"`
	Status           int    `gorm:"type:int;not null" json:"status"`
	UserID           string `gorm:"type:varchar(191);not null" json:"user_id"`
	// UserEntity       UserEntity `gorm:"references:IDUser"`
}

func (SuratPerjalananDinasEntity) TableName() string {
	return "surat_perjalanan_dinas"
}
