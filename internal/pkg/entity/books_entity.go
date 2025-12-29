package entity

import (
	"time"

	"gorm.io/gorm"
)

type (
	Book struct {
		gorm.Model
		UserID      uint
		User        User
		Title       string
		Description string
		Author      string
	}

	FilterBooks struct {
		Limit, Offset int
		Title         string
	}
)

// User model
type UserX struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	NamaUser      string    `gorm:"size:255;not null;index"`
	KataSandi     string    `gorm:"size:255;not null"`
	NoTelp        string    `gorm:"size:255;unique;not null;index"`
	TanggalLahir  time.Time `gorm:"not null"`
	JenisKelamin  string    `gorm:"size:255"`
	TentangText   string    `gorm:"type:text"`
	Pekerjaan     string    `gorm:"size:255"`
	Email         string    `gorm:"size:255;not null;index"`
	IDProvinsi    string    `gorm:"size:255"`
	IDKota        string    `gorm:"size:255"`
	IsAdmin       bool      `gorm:"default:false"`
	UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
	CreatedAtDate time.Time `gorm:"autoCreateTime"`
}

// Alamat model
type AlamatX struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	IDUser        uint      `gorm:"not null;index"` // Index untuk pencarian cepat berdasarkan IDUser
	JudulAlamat   string    `gorm:"size:255;not null"`
	NamaPenerima  string    `gorm:"size:255;not null"`
	NoTelp        string    `gorm:"size:255;not null"`
	DetailAlamat  string    `gorm:"size:255;not null"`
	UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
	CreatedAtDate time.Time `gorm:"autoCreateTime"`
	User          User      `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE"`
}

// Toko model
type TokoX struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	IDUser        uint      `gorm:"not null;index"` // Index untuk pencarian cepat berdasarkan IDUser
	NamaToko      string    `gorm:"size:255;not null"`
	URLFoto       string    `gorm:"size:255"`
	UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
	CreatedAtDate time.Time `gorm:"autoCreateTime"`
	User          User      `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE"`
}

// Category model
type CategoryX struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	NamaCategory  string    `gorm:"size:255;not null;index"` // Index pada kategori
	CreatedAtDate time.Time `gorm:"autoCreateTime"`
	UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
}

// Produk model
type ProdukX struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	NamaProduk    string    `gorm:"size:255;not null;index"`        // Index pada Nama Produk
	Slug          string    `gorm:"size:255;unique;not null;index"` // Unique index pada Slug
	HargaReseller string    `gorm:"size:255;not null"`
	HargaKonsumen string    `gorm:"size:255;not null"`
	Stok          int       `gorm:"default:0"`
	Deskripsi     string    `gorm:"type:text"`
	CreatedAtDate time.Time `gorm:"autoCreateTime"`
	UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
	IDCategory    uint      `gorm:"not null;index"` // Index untuk pencarian berdasarkan kategori
	Category      Category  `gorm:"foreignKey:IDCategory;constraint:OnDelete:CASCADE"`
}

// FotoProduk model
type FotoProdukX struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	IDProduk      uint      `gorm:"not null;index"` // Index pada IDProduk untuk mempercepat pencarian foto produk
	URL           string    `gorm:"size:255;not null"`
	UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
	CreatedAtDate time.Time `gorm:"autoCreateTime"`
	Produk        Produk    `gorm:"foreignKey:IDProduk;constraint:OnDelete:CASCADE"`
}

// LogProduk model
type LogProdukX struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	IDProduk      uint      `gorm:"not null;index"` // Index pada IDProduk untuk pencarian log produk
	NamaProduk    string    `gorm:"size:255;not null"`
	Slug          string    `gorm:"size:255;not null"`
	HargaReseller string    `gorm:"size:255;not null"`
	HargaKonsumen string    `gorm:"size:255;not null"`
	Stok          int       `gorm:"default:0"`
	Deskripsi     string    `gorm:"type:text"`
	CreatedAtDate time.Time `gorm:"autoCreateTime"`
	UpdatedAtDate time.Time `gorm:"autoUpdateTime"`
	IDToko        uint      `gorm:"not null;index"` // Index pada IDToko untuk pencarian log berdasarkan toko
	IDCategory    uint      `gorm:"not null;index"` // Index pada IDCategory untuk pencarian berdasarkan kategori
	Toko          Toko      `gorm:"foreignKey:IDToko;constraint:OnDelete:CASCADE"`
	Category      Category  `gorm:"foreignKey:IDCategory;constraint:OnDelete:CASCADE"`
	Produk        Produk    `gorm:"foreignKey:IDProduk;constraint:OnDelete:CASCADE"`
}

// Transaction model
type TransactionX struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	IDUser           uint      `gorm:"not null;index"` // Index untuk pencarian berdasarkan IDUser
	AlamatPengiriman uint      `gorm:"not null;index"` // Index pada AlamatPengiriman untuk optimasi pencarian pengiriman
	HargaTotal       int       `gorm:"not null"`
	KodeInvoice      string    `gorm:"size:255;not null;index"` // Index pada KodeInvoice untuk pencarian berdasarkan invoice
	MethodBayar      string    `gorm:"size:255;not null"`
	UpdatedAtDate    time.Time `gorm:"autoUpdateTime"`
	CreatedAtDate    time.Time `gorm:"autoCreateTime"`
	User             User      `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE"`
	Alamat           Alamat    `gorm:"foreignKey:AlamatPengiriman;constraint:OnDelete:CASCADE"`
}

// DetailTransaction model
type DetailTransactionX struct {
	ID            uint        `gorm:"primaryKey;autoIncrement"`
	IDTrx         uint        `gorm:"not null;index"` // Index untuk pencarian berdasarkan IDTransaksi
	IDProduk      uint        `gorm:"not null;index"` // Index pada IDProduk untuk pencarian berdasarkan produk
	Kuantitas     int         `gorm:"not null"`
	HargaTotal    int         `gorm:"not null"`
	UpdatedAtDate time.Time   `gorm:"autoUpdateTime"`
	CreatedAtDate time.Time   `gorm:"autoCreateTime"`
	Transaction   Transaction `gorm:"foreignKey:IDTrx;constraint:OnDelete:CASCADE"`
	Produk        Produk      `gorm:"foreignKey:IDProduk;constraint:OnDelete:CASCADE"`
}
