package entity

import "time"

type (
	Toko struct {
		ID            uint      `gorm:"primaryKey;autoIncrement"`
		IDUser        uint      `gorm:"not null;index"`
		NamaToko      string    `gorm:"size:255;not null"`
		URLFoto       string    `gorm:"size:255"`
		UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
		CreatedAtDate time.Time `gorm:"autoCreateTime"`

		User User `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE"`
	}

	FilterToko struct {
		Limit, Offset int
		Name          string
	}
)
