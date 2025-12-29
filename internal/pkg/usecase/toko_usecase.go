package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	tokoModel "tugas_akhir_example/internal/pkg/model"
	tokoRepository "tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TokoUseCase interface {
	GetAllToko(ctx context.Context, params tokoModel.TokosFilter) (res []tokoModel.TokosResp, err *helper.ErrorStruct)
	GetTokoByID(ctx context.Context, booksid string) (res tokoModel.TokosResp, err *helper.ErrorStruct)
	CreateToko(ctx context.Context, data entity.Toko) (res uint, err *helper.ErrorStruct)
	UpdateTokoByID(ctx context.Context, booksid string, data tokoModel.TokosReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteTokoByID(ctx context.Context, booksid string, userId string) (res string, err *helper.ErrorStruct)
}

type TokoUseCaseImpl struct {
	tokoRepository tokoRepository.TokoRepository
}

func NewTokoUseCase(tokoRepository tokoRepository.TokoRepository) TokoUseCase {
	return &TokoUseCaseImpl{
		tokoRepository: tokoRepository,
	}

}

func (alc *TokoUseCaseImpl) GetAllToko(ctx context.Context, params tokoModel.TokosFilter) (res []tokoModel.TokosResp, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := alc.tokoRepository.GetAllToko(ctx, entity.FilterToko{
		Limit:  params.Limit,
		Offset: params.Page,
		Name:   params.Name,
	})
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Toko"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllToko : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, tokoModel.TokosResp{
			ID:       v.ID,
			NamaToko: v.NamaToko,
			URLFoto:  v.URLFoto,
		})
	}

	return res, nil
}
func (alc *TokoUseCaseImpl) GetTokoByID(ctx context.Context, booksid string) (res tokoModel.TokosResp, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.tokoRepository.GetTokoByID(ctx, booksid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Toko"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetTokoByID : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = tokoModel.TokosResp{
		ID:       resRepo.ID,
		NamaToko: resRepo.NamaToko,
		URLFoto:  resRepo.URLFoto,
	}

	return res, nil
}
func (alc *TokoUseCaseImpl) CreateToko(ctx context.Context, data entity.Toko) (res uint, err *helper.ErrorStruct) {

	resRepo, errRepo := alc.tokoRepository.CreateToko(ctx, data)
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateToko : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *TokoUseCaseImpl) UpdateTokoByID(ctx context.Context, booksid string, data tokoModel.TokosReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.tokoRepository.UpdateTokoByID(ctx, booksid, entity.Toko{
		NamaToko: data.NamaToko,
		URLFoto:  "MASIH HARDCODE",
	})

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at UpdateTokoByID : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *TokoUseCaseImpl) DeleteTokoByID(ctx context.Context, booksid string, userId string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.tokoRepository.DeleteTokoByID(ctx, booksid, userId)
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at DeleteTokoByID : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
