package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	categoriesModel "tugas_akhir_example/internal/pkg/model"
	categoriesRepositories "tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoriesUseCase interface {
	GetAllCategories(ctx context.Context) (res []categoriesModel.CategoriesRes, err *helper.ErrorStruct)
	GetCategoriesByID(ctx context.Context, categoryId string) (res categoriesModel.CategoriesRes, err *helper.ErrorStruct)
	CreateCategories(ctx context.Context, data categoriesModel.CategoriesReq) (res uint, err *helper.ErrorStruct)
	UpdateCategoriesByID(ctx context.Context, categoryId string, data categoriesModel.CategoriesReq) (res string, err *helper.ErrorStruct)
	DeleteCategoriesByID(ctx context.Context, categoryId string) (res string, err *helper.ErrorStruct)
}

type CategoriesUseCaseImpl struct {
	categoriesRepositories categoriesRepositories.CategoriesRepository
}

func NewCategoriesUseCase(categoriesRepositories categoriesRepositories.CategoriesRepository) CategoriesUseCase {
	return &CategoriesUseCaseImpl{
		categoriesRepositories: categoriesRepositories,
	}

}

func (alc *CategoriesUseCaseImpl) GetAllCategories(ctx context.Context) (res []categoriesModel.CategoriesRes, err *helper.ErrorStruct) {

	resRepo, errRepo := alc.categoriesRepositories.GetAllCategories(ctx)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Category"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllCategories : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, categoriesModel.CategoriesRes{
			ID:           v.ID,
			NamaCategory: v.NamaCategory,
		})
	}

	return res, nil
}
func (alc *CategoriesUseCaseImpl) GetCategoriesByID(ctx context.Context, categoryId string) (res categoriesModel.CategoriesRes, err *helper.ErrorStruct) {

	resRepo, errRepo := alc.categoriesRepositories.GetCategoriesByID(ctx, categoryId)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Category"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetCategoriesByID : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = categoriesModel.CategoriesRes{
		ID:           resRepo.ID,
		NamaCategory: resRepo.NamaCategory,
	}

	return res, nil
}
func (alc *CategoriesUseCaseImpl) CreateCategories(ctx context.Context, data categoriesModel.CategoriesReq) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.categoriesRepositories.CreateCategories(ctx, entity.Category{
		NamaCategory: data.NamaCategory,
	})
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateCategories : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *CategoriesUseCaseImpl) UpdateCategoriesByID(ctx context.Context, categoryId string, data categoriesModel.CategoriesReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.categoriesRepositories.UpdateCategoriesByID(ctx, categoryId, entity.Category{
		NamaCategory: data.NamaCategory,
	})

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at UpdateCategoriesByID : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *CategoriesUseCaseImpl) DeleteCategoriesByID(ctx context.Context, categoryId string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.categoriesRepositories.DeleteCategoriesByID(ctx, categoryId)
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at DeleteCategoriesByID : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
