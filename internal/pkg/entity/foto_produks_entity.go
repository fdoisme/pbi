package entity

import "time"

type (
	FotoProduk struct {
		ID            uint      `gorm:"primaryKey;autoIncrement"`
		IDProduk      uint      `gorm:"not null;index"`
		URL           string    `gorm:"size:255;not null"`
		UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
		CreatedAtDate time.Time `gorm:"autoCreateTime"`

		// Produk Produk `gorm:"foreignKey:IDProduk;constraint:OnDelete:CASCADE"`
	}

	FilterFotoProduk struct {
		Limit, Offset int
		Title         string
	}
)
