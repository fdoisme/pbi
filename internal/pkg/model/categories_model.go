package model

type CategoriesReq struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}
type CategoriesRes struct {
	ID           uint   `json:"id"`
	NamaCategory string `json:"nama_category"`
}
