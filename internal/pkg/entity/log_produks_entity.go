package entity

import "time"

type (
	LogProduk struct {
		ID            uint      `gorm:"primaryKey;autoIncrement"`
		IDProduk      uint      `gorm:"not null;index"`
		NamaProduk    string    `gorm:"size:255;not null"`
		Slug          string    `gorm:"size:255;not null"`
		HargaReseller string    `gorm:"size:255;not null"`
		HargaKonsumen string    `gorm:"size:255;not null"`
		Stok          int       `gorm:"default:0"`
		Deskripsi     string    `gorm:"type:text"`
		CreatedAtDate time.Time `gorm:"autoCreateTime"`
		UpdatedAtDate time.Time `gorm:"autoUpdateTime"`

		IDToko     uint `gorm:"not null;index"`
		IDCategory uint `gorm:"not null;index"`

		Toko     Toko     `gorm:"foreignKey:IDToko;constraint:OnDelete:CASCADE"`
		Category Category `gorm:"foreignKey:IDCategory;constraint:OnDelete:CASCADE"`
		Produk   Produk   `gorm:"foreignKey:IDProduk;constraint:OnDelete:CASCADE"`
	}

	FilterLogProduk struct {
		Limit, Offset int
		Title         string
	}
)
