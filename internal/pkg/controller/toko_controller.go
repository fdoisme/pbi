package controller

import (
	"errors"
	"log"
	"tugas_akhir_example/internal/helper"
	tokoModel "tugas_akhir_example/internal/pkg/model"
	tokoUseCase "tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type TokoController interface {
	GetAllToko(ctx *fiber.Ctx) error
	GetTokoByID(ctx *fiber.Ctx) error
	UpdateTokoByID(ctx *fiber.Ctx) error
	DeleteTokoByID(ctx *fiber.Ctx) error
}

type TokoControllerImpl struct {
	tokoUseCase tokoUseCase.TokoUseCase
}

func NewTokoController(tokoUseCase tokoUseCase.TokoUseCase) TokoController {
	return &TokoControllerImpl{
		tokoUseCase: tokoUseCase,
	}
}

func (uc *TokoControllerImpl) GetAllToko(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(tokoModel.TokosFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request").Error())
		return nil
	}

	res, err := uc.tokoUseCase.GetAllToko(c, tokoModel.TokosFilter{
		Name:  filter.Name,
		Limit: filter.Limit,
		Page:  filter.Page,
	})
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Err.Error())
		return nil
	}
	pagination := helper.DataWithPagination[[]tokoModel.TokosResp]{
		Page:  filter.Page,
		Limit: filter.Limit,
		Data:  res,
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", pagination)
	return nil
}

func (uc *TokoControllerImpl) GetTokoByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	tokoId := ctx.Params("id_toko")
	if tokoId == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request").Error())
		return nil
	}

	res, err := uc.tokoUseCase.GetTokoByID(c, tokoId)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Err.Error())
		return nil
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", res)
	return nil
}

func (uc *TokoControllerImpl) UpdateTokoByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	tokoId := ctx.Params("id_toko")
	if tokoId == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request").Error())
		return nil
	}

	data := new(tokoModel.TokosReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request").Error())
		return nil
	}

	res, err := uc.tokoUseCase.UpdateTokoByID(c, tokoId, *data)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", err.Err.Error())
		return nil
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", res)
	return nil
}

func (uc *TokoControllerImpl) DeleteTokoByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	tokoId := ctx.Params("id_toko")
	userID := ctx.Locals("userid").(string)
	if tokoId == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, fiber.StatusBadRequest, "Failed to GET data", errors.New("Bad request").Error())
		return nil
	}

	res, err := uc.tokoUseCase.DeleteTokoByID(c, tokoId, userID)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		helper.BadResponse[any](ctx, err.Code, "Failed to GET data", err.Err.Error())
		return err.Err
	}

	// TODO IMRPOVE FORMAT RESPONSE
	helper.GoodResponse(ctx, "Succeed to GET data", res)
	return nil
}
