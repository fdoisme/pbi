package entity

import "time"

type (
	Alamat struct {
		ID            uint      `gorm:"primaryKey;autoIncrement"`
		IDUser        uint      `gorm:"not null;index"`
		JudulAlamat   string    `gorm:"size:255;not null"`
		NamaPenerima  string    `gorm:"size:255;not null"`
		NoTelp        string    `gorm:"size:255;not null"`
		DetailAlamat  string    `gorm:"size:255;not null"`
		UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
		CreatedAtDate time.Time `gorm:"autoCreateTime"`

		User User `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE"`
	}

	FilteAlamat struct {
		Limit, Offset int
		Title         string
	}
)
