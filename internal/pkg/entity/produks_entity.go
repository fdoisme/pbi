package entity

import "time"

type (
	Produk struct {
		ID            uint         `gorm:"primaryKey;autoIncrement"`
		NamaProduk    string       `gorm:"size:255;not null;index"`
		Slug          string       `gorm:"size:255;not null;index"`
		HargaReseller int          `gorm:"not null"`
		HargaKonsumen int          `gorm:"not null"`
		Stok          int          `gorm:"default:0"`
		Deskripsi     string       `gorm:"type:text"`
		CreatedAtDate time.Time    `gorm:"autoCreateTime"`
		UpdatedAtDate time.Time    `gorm:"autoUpdateTime"`
		IDToko        uint         `gorm:"not null;index"`
		Toko          Toko         `gorm:"foreignKey:IDToko;constraint:OnDelete:CASCADE"`
		IDCategory    uint         `gorm:"not null;index"`
		Category      Category     `gorm:"foreignKey:IDCategory;constraint:OnDelete:CASCADE"`
		Foto          []FotoProduk `gorm:"foreignKey:IDProduk;constraint:OnDelete:CASCADE"`
	}

	FilterProduk struct {
		Limit, Offset int
		NamaProduk    string
		CategoryID    string
		TokoID        string
		MaxHarga      string
		MinHarga      string
	}
)
