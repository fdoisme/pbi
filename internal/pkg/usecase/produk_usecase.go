package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	produkModel "tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"
	"tugas_akhir_example/internal/utils/mapper"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProdukUseCase interface {
	GetAllProduk(ctx context.Context, params *produkModel.ProdukFilter) (res []produkModel.ProdukRes, err *helper.ErrorStruct)
	GetProdukByID(ctx context.Context, produkId string) (res model.ProdukRes, err *helper.ErrorStruct)
	CreateProduk(ctx context.Context, userid string, data produkModel.ProdukRequest, urls []string) (res uint, err *helper.ErrorStruct)
	UpdateProdukByID(ctx context.Context, userid string, produkId string, data produkModel.ProdukRequest, urls []string) (res string, err *helper.ErrorStruct)
	DeleteProdukByID(ctx context.Context, produkId string, userId string) (res string, err *helper.ErrorStruct)
}

type ProdukUseCaseImpl struct {
	produkRepository     repository.ProdukRepository
	tokoRepository       repository.TokoRepository
	fotoProdukRepository repository.FotoProdukRepository
}

func NewProdukUseCase(produkRepository repository.ProdukRepository, tokoRepository repository.TokoRepository, fotoProdukRepository repository.FotoProdukRepository) ProdukUseCase {
	return &ProdukUseCaseImpl{
		produkRepository:     produkRepository,
		tokoRepository:       tokoRepository,
		fotoProdukRepository: fotoProdukRepository,
	}

}

func (alc *ProdukUseCaseImpl) GetAllProduk(ctx context.Context, params *produkModel.ProdukFilter) (res []produkModel.ProdukRes, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		fmt.Println("[LIMIT] MASUK 0:", params.Limit)
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}
	fmt.Println("[FILTER]:", params)

	resRepo, errRepo := alc.produkRepository.GetAllProduk(ctx, entity.FilterProduk{
		Limit:      params.Limit,
		Offset:     params.Page,
		NamaProduk: params.NamaProduk,
		CategoryID: params.CategoryID,
		TokoID:     params.TokoID,
		MinHarga:   params.MinHarga,
		MaxHarga:   params.MaxHarga,
	})
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Produk"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllProduk : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	jsonData, _ := json.Marshal(resRepo)

	// Mencetak JSON
	fmt.Println(string(jsonData))
	for _, v := range resRepo {
		res = append(res, mapper.MapperToProdukRespon(v))
	}
	return res, nil
}
func (alc *ProdukUseCaseImpl) GetProdukByID(ctx context.Context, produkId string) (res produkModel.ProdukRes, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.produkRepository.GetProdukByID(ctx, produkId)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Produk"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllProduk : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	jsonData, _ := json.Marshal(resRepo)
	fmt.Println(string(jsonData))
	// res = produkModel.ProdukRes{
	// 	ID:          resRepo.ID,
	// 	Title:       resRepo.Title,
	// 	Description: resRepo.Description,
	// 	Author:      resRepo.Author,
	// }
	res = mapper.MapperToProdukRespon(resRepo)
	return res, nil
}

func (alc *ProdukUseCaseImpl) CreateProduk(ctx context.Context, userid string, data produkModel.ProdukRequest, urls []string) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}
	toko, errRepo := alc.tokoRepository.GetTokoByUserID(ctx, userid)
	if errRepo != nil {
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	categoryID := utils.StrToUint(data.CategoryID)
	hargaReseller := utils.StrToInt(data.HargaReseller)
	hargaKonsumen := utils.StrToInt(data.HargaKonsumen)
	stock := utils.StrToInt(data.Stok)
	resRepo, errRepo := alc.produkRepository.CreateProduk(ctx, entity.Produk{
		NamaProduk:    data.NamaProduk,
		IDCategory:    uint(categoryID),
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stock,
		Deskripsi:     data.Deskripsi,
		IDToko:        toko.ID,
		Slug:          strings.Replace(data.NamaProduk, " ", "-", -1),
	})
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduk : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	for _, v := range urls {
		_, errRepo = alc.fotoProdukRepository.CreateProduk(ctx, entity.FotoProduk{URL: v, IDProduk: resRepo})
		if errRepo != nil {
			helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduk : %s", errRepo.Error()), errRepo)
			return res, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  errRepo,
			}
		}
	}
	return resRepo, nil
}

func (alc *ProdukUseCaseImpl) UpdateProdukByID(ctx context.Context, userid string, produkId string, data produkModel.ProdukRequest, urls []string) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}
	produk, errRepo := alc.produkRepository.GetProdukByID(ctx, produkId)
	if errRepo != nil {
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	fmt.Println("[APAKAH BISA AKSES] : ", userid, produk)
	if userid != fmt.Sprint(produk.Toko.IDUser) {
		err := errors.New("No Authorization")
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at UpdateProduk : %s", err.Error()), err)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}
	categoryID := utils.StrToUint(data.CategoryID)
	hargaReseller := utils.StrToInt(data.HargaReseller)
	hargaKonsumen := utils.StrToInt(data.HargaKonsumen)
	stock := utils.StrToInt(data.Stok)
	resRepo, errRepo := alc.produkRepository.UpdateProdukByID(ctx, produkId, entity.Produk{
		NamaProduk:    data.NamaProduk,
		IDCategory:    uint(categoryID),
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stock,
		Deskripsi:     data.Deskripsi,
		IDToko:        produk.IDToko,
		Slug:          strings.ReplaceAll(data.NamaProduk, " ", "-"),
	})
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at UpdateProduk : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	resRepo, errRepo = alc.fotoProdukRepository.UpdateProdukByProdukID(ctx, produkId, urls)
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at UpdateProduk : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	return resRepo, nil
}

func (alc *ProdukUseCaseImpl) DeleteProdukByID(ctx context.Context, produkId string, userId string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.produkRepository.DeleteProdukByID(ctx, produkId, userId)
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllProduk : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
