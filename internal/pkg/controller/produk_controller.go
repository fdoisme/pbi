package controller

import (
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/helper"
	produkmodel "tugas_akhir_example/internal/pkg/model"
	produkUseCase "tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProdukController interface {
	GetAllProduk(ctx *fiber.Ctx) error
	GetProdukByID(ctx *fiber.Ctx) error
	CreateProduk(ctx *fiber.Ctx) error
	UpdateProdukByID(ctx *fiber.Ctx) error
	DeleteProdukByID(ctx *fiber.Ctx) error
}

type ProdukControllerImpl struct {
	produkUseCase produkUseCase.ProdukUseCase
}

func NewProdukController(produkUseCase produkUseCase.ProdukUseCase) ProdukController {
	return &ProdukControllerImpl{
		produkUseCase: produkUseCase,
	}
}

func (uc *ProdukControllerImpl) GetAllProduk(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(produkmodel.ProdukFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := uc.produkUseCase.GetAllProduk(c, filter)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	pagination := helper.DataWithPagination[[]produkmodel.ProdukRes]{
		Page:  filter.Page,
		Limit: filter.Limit,
		Data:  res,
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", pagination)
	return nil
}

func (uc *ProdukControllerImpl) GetProdukByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	produkId := ctx.Params("id_produk")
	if produkId == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	res, err := uc.produkUseCase.GetProdukByID(c, produkId)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *ProdukControllerImpl) CreateProduk(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid").(string)

	data := new(produkmodel.ProdukRequest)
	if err := ctx.BodyParser(data); err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Error())
		return nil
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Error())
		return nil
	}
	files := form.File["photos"]
	if len(files) == 0 {
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request (Photos)").Error())
		return nil
	}
	urls, err := utils.UploadFiles(ctx, files)
	if err != nil {
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Error())
		return nil
	}
	fmt.Println("[URL]", urls)
	res, errUseCase := uc.produkUseCase.CreateProduk(c, userid, *data, urls)
	if errUseCase != nil {
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errUseCase.Err.Error())
		return nil
	}
	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *ProdukControllerImpl) UpdateProdukByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	produkId := ctx.Params("id_produk")
	userid := ctx.Locals("userid").(string)

	data := new(produkmodel.ProdukRequest)
	if err := ctx.BodyParser(data); err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Error())
		return nil
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Error())
		return nil
	}
	files := form.File["photos"]
	if len(files) == 0 {
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request (Photos)").Error())
		return nil
	}
	urls, err := utils.UploadFiles(ctx, files)
	if err != nil {
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Error())
		return nil
	}
	fmt.Println("[URL]", urls)
	res, errUseCase := uc.produkUseCase.UpdateProdukByID(c, userid, produkId, *data, urls)
	if errUseCase != nil {
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errUseCase.Err.Error())
		return nil
	}
	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *ProdukControllerImpl) DeleteProdukByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	produkId := ctx.Params("id_produk")
	userID := ctx.Locals("userid").(string)
	if produkId == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	res, err := uc.produkUseCase.DeleteProdukByID(c, produkId, userID)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}
