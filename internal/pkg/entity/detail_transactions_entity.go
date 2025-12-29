package entity

import "time"

type (
	DetailTransaction struct {
		ID            uint      `gorm:"primaryKey;autoIncrement"`
		IDTrx         uint      `gorm:"not null;index"`
		IDProduk      uint      `gorm:"not null;index"`
		Kuantitas     int       `gorm:"not null"`
		HargaTotal    int       `gorm:"not null"`
		UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
		CreatedAtDate time.Time `gorm:"autoCreateTime"`

		Transaction Transaction `gorm:"foreignKey:IDTrx;constraint:OnDelete:CASCADE"`
		Produk      Produk      `gorm:"foreignKey:IDProduk;constraint:OnDelete:CASCADE"`
	}

	FilterDetailTransaction struct {
		Limit, Offset int
		Title         string
	}
)
