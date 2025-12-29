package model

import "github.com/gofiber/fiber/v2"

type TokosFilter struct {
	Name  string `query:"nama"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
}

type TokosReqUpdate struct {
	NamaToko string          `form:"NamaToko,omitempty"`
	Photo    *fiber.FormFile `form:"photo,omitempty"`
}

type MyTokoResp struct {
	ID       uint   `json:"id"`
	IDUser   uint   `json:"user_id"`
	NamaToko string `json:"nama_toko"`
	URLFoto  string `json:"url_foto"`
}

type TokosResp struct {
	ID       uint   `json:"id"`
	NamaToko string `json:"nama_toko"`
	URLFoto  string `json:"url_foto"`
}
