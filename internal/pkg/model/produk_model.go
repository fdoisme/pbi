package model

type ProdukFilter struct {
	Limit      int    `query:"limit"`
	Page       int    `query:"page"`
	NamaProduk string `query:"nama_produk"`
	CategoryID string `query:"category_id"`
	TokoID     string `query:"toko_id"`
	MaxHarga   string `query:"max_harga"`
	MinHarga   string `query:"min_harga"`
}

type ProdukRes struct {
	ID            uint          `json:"id"`
	NamaProduk    string        `json:"nama_produk"`
	Slug          string        `json:"slug"`
	HargaReseller int           `json:"harga_reseller"`
	HargaKonsumen int           `json:"harga_konsumen"`
	Stok          int           `json:"stok"`
	Deskripsi     string        `json:"deskripsi"`
	Toko          TokosResp     `json:"toko"`
	Category      CategoriesRes `json:"category"`
	Photos        []FotoProdukRes `json:"photos"`
}

type ProdukRequest struct {
	NamaProduk    string `form:"nama_produk"`
	CategoryID    string `form:"category_id"`
	HargaReseller string `form:"harga_reseller"`
	HargaKonsumen string `form:"harga_konsumen"`
	Stok          string `form:"stok"`
	Deskripsi     string `form:"deskripsi"`
}
