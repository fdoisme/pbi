package model

type FotoProdukReq struct {
	Url string
}

type FotoProdukRes struct {
	ID       uint   `json:"id"`
	IDProduk uint   `json:"product_id"`
	URL      string `json:"url"`
}
