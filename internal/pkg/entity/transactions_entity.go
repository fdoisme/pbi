package entity

import "time"

type (
	Transaction struct {
		ID               uint      `gorm:"primaryKey;autoIncrement"`
		IDUser           uint      `gorm:"not null;index"`
		AlamatPengiriman uint      `gorm:"not null;index"`
		HargaTotal       int       `gorm:"not null"`
		KodeInvoice      string    `gorm:"size:255;not null;index"`
		MethodBayar      string    `gorm:"size:255;not null"`
		UpdatedAtDate    time.Time `gorm:"autoUpdateTime"`
		CreatedAtDate    time.Time `gorm:"autoCreateTime"`

		User   User   `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE"`
		Alamat Alamat `gorm:"foreignKey:AlamatPengiriman;constraint:OnDelete:CASCADE"`
	}

	FilterTransaction struct {
		Limit, Offset int
		Title         string
	}
)
