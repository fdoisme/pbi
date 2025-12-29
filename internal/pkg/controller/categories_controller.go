package controller

import (
	"errors"
	"fmt"
	"tugas_akhir_example/internal/helper"
	categoriesModel "tugas_akhir_example/internal/pkg/model"
	categoriesUseCase "tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoriesController interface {
	GetAllCategories(ctx *fiber.Ctx) error
	GetCategoriesByID(ctx *fiber.Ctx) error
	CreateCategories(ctx *fiber.Ctx) error
	UpdateCategoriesByID(ctx *fiber.Ctx) error
	DeleteCategoriesByID(ctx *fiber.Ctx) error
}

type CategoriesControllerImpl struct {
	categoriesUseCase categoriesUseCase.CategoriesUseCase
}

func NewCategoriesController(categoriesUseCase categoriesUseCase.CategoriesUseCase) CategoriesController {
	return &CategoriesControllerImpl{
		categoriesUseCase: categoriesUseCase,
	}
}

func (uc *CategoriesControllerImpl) GetAllCategories(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, err := uc.categoriesUseCase.GetAllCategories(c)

	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, err.Code, "Failed to GET data", err.Err.Error())
		return nil
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", res)
	return nil
}

func (uc *CategoriesControllerImpl) GetCategoriesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId := ctx.Params("id_category")
	if categoryId == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request").Error())
		return nil
	}

	res, err := uc.categoriesUseCase.GetCategoriesByID(c, categoryId)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, err.Code, "Failed to GET data", err.Err.Error())
		return nil
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", res)
	return nil
}

func (uc *CategoriesControllerImpl) CreateCategories(ctx *fiber.Ctx) error {
	c := ctx.Context()

	// cara baca context yang diset di middleware
	id := ctx.Locals("userid").(string)
	email := ctx.Locals("useremail").(string)

	fmt.Println("id", id)
	fmt.Println("email", email)

	data := new(categoriesModel.CategoriesReq)
	if err := ctx.BodyParser(data); err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Error())
		return nil
	}
	res, err := uc.categoriesUseCase.CreateCategories(c, *data)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Err.Error())
		return nil
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", res)
	return nil
}

func (uc *CategoriesControllerImpl) UpdateCategoriesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId := ctx.Params("id_category")
	if categoryId == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request").Error())
		return nil
	}

	data := new(categoriesModel.CategoriesReq)
	if err := ctx.BodyParser(data); err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Error())
		return nil
	}

	res, err := uc.categoriesUseCase.UpdateCategoriesByID(c, categoryId, *data)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Err.Error())
		return nil
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", res)
	return nil
}

func (uc *CategoriesControllerImpl) DeleteCategoriesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId := ctx.Params("id_category")
	if categoryId == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request").Error())
		return nil
	}

	res, err := uc.categoriesUseCase.DeleteCategoriesByID(c, categoryId)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Err.Error())
		return nil
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", res)
	return nil
}
