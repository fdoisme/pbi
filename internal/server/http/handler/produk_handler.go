package handler

import (
	"github.com/gofiber/fiber/v2"

	produkController "tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
)

func ProdukRoute(r fiber.Router, ProdukUsc usecase.ProdukUseCase) {
	controller := produkController.NewProdukController(ProdukUsc)

	produkAPI := r.Group("/product")
	produkAPI.Get("/", controller.GetAllProduk)
	produkAPI.Get("/:id_produk", controller.GetProdukByID)
	produkAPI.Use(MiddlewareAuth)
	produkAPI.Post("", controller.CreateProduk)
	produkAPI.Put("/:id_produk", controller.UpdateProdukByID)
	produkAPI.Delete("/:id_produk", controller.DeleteProdukByID)
}
