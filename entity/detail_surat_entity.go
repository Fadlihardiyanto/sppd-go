package entity

type DetailSuratEntity struct {
	IDDetailSurat string `gorm:"primaryKey;autoIncrement;not null" json:"detail_surat_id"`
	SuratID       string `gorm:"type:varchar(255);not null" json:"surat_id"`
	// SuratPerjalananDinasEntity SuratPerjalananDinasEntity `gorm:"references:IDSurat"`
	Deskripsi string `gorm:"type:varchar(255);not null" json:"deskripsi"`
}

func (DetailSuratEntity) TableName() string {
	return "detail_surat"
}
