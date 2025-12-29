package entity

import "time"

type (
	User struct {
		ID            uint   `gorm:"primaryKey;autoIncrement"`
		NamaUser      string `gorm:"size:255"`
		KataSandi     string `gorm:"size:255"`
		NoTelp        string `gorm:"size:255;unique"`
		TanggalLahir  string `gorm:"type:date"`
		JenisKelamin  string `gorm:"size:255"`
		TentangText   string `gorm:"type:text"`
		Pekerjaan     string `gorm:"size:255"`
		Email         string `gorm:"size:255"`
		IDProvinsi    string `gorm:"size:255"`
		IDKota        string `gorm:"size:255"`
		IsAdmin       bool
		UpdatedAtDate time.Time `gorm:"type:timestamp;default:now()"`
		CreatedAtDate time.Time `gorm:"type:timestamp;default:now()"`
	}

	FilterUser struct {
		Limit, Offset int
		Title         string
	}
)
