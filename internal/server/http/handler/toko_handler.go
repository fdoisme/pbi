package handler

import (
	"github.com/gofiber/fiber/v2"

	tokoController "tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
)

func TokoRoute(r fiber.Router, TokoUsc usecase.TokoUseCase) {
	controller := tokoController.NewTokoController(TokoUsc)

	tokoAPI := r.Group("/toko")
	tokoAPI.Get("", controller.GetAllToko)
	tokoAPI.Get("/my", controller.GetTokoByID)
	tokoAPI.Get("/:id_toko", controller.GetTokoByID)
	tokoAPI.Use(MiddlewareAuth)
	tokoAPI.Put("/:id_toko", controller.UpdateTokoByID)
	tokoAPI.Delete("/:id_toko", controller.DeleteTokoByID)
}
