package handler

import (
	"github.com/gofiber/fiber/v2"

	categoriescontroller "tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
)

func CategoriesRoute(r fiber.Router, CategoriesUsc usecase.CategoriesUseCase) {
	controller := categoriescontroller.NewCategoriesController(CategoriesUsc)

	categoryAPI := r.Group("/category")
	categoryAPI.Get("", controller.GetAllCategories)
	categoryAPI.Get("/:id_category", controller.GetCategoriesByID)
	categoryAPI.Use(MiddlewareAuth)
	categoryAPI.Post("", controller.CreateCategories)
	categoryAPI.Put("/:id_category", controller.UpdateCategoriesByID)
	categoryAPI.Delete("/:id_category", controller.DeleteCategoriesByID)
}
