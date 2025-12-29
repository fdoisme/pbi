package entity

import "time"

type (
	Category struct {
		ID            uint      `gorm:"primaryKey;autoIncrement"`
		NamaCategory  string    `gorm:"size:255;not null;index"`
		CreatedAtDate time.Time `gorm:"autoCreateTime"`
		UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
	}

	FilterCategory struct {
		Limit, Offset int
		Title         string
	}
)
