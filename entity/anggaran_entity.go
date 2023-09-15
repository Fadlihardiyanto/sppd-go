package entity

type KategoryAnggaranEntity struct {
	IDKategory   string `gorm:"primaryKey;not null" json:"category_id"`
	NamaKategori string `gorm:"type:varchar(255);not null" json:"nama_kategori"`
}

func (KategoryAnggaranEntity) TableName() string {
	return "kategori_anggaran"
}
