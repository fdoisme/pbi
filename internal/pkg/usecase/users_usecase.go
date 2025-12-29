package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	usermodel "tugas_akhir_example/internal/pkg/model"
	userrepository "tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UsersUseCase interface {
	Login(ctx context.Context, params usermodel.Login) (res usermodel.LoginRes, err *helper.ErrorStruct)
	CreateUsers(ctx context.Context, data usermodel.CreateUser) (res uint, err *helper.ErrorStruct)
}

type UsersUseCaseImpl struct {
	userrepository userrepository.UsersRepository
}

func NewUsersUseCase(userrepository userrepository.UsersRepository) UsersUseCase {
	return &UsersUseCaseImpl{
		userrepository: userrepository,
	}

}

func (alc *UsersUseCaseImpl) Login(ctx context.Context, params usermodel.Login) (res usermodel.LoginRes, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.userrepository.GetUsersByNoTelp(ctx, params.NoTelp)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Users"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllUsers : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	// fmt.Println("[USECASE]", params.KataSandi, hashPass)
	isValid := utils.CheckPasswordHash(params.KataSandi, resRepo.KataSandi)
	if !isValid {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("invalid account"),
		}
	}

	tokenInit := utils.NewToken(utils.DataClaims{
		ID:      fmt.Sprint(resRepo.ID),
		Email:   resRepo.Email,
		IsAdmin: resRepo.IsAdmin,
	})

	token, errToken := tokenInit.Create()
	if errToken != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errToken,
		}
	}

	res = usermodel.LoginRes{
		Email: resRepo.Email,
		Name:  resRepo.NamaUser,
		Token: token,
	}

	return res, nil
}
func (alc *UsersUseCaseImpl) CreateUsers(ctx context.Context, params usermodel.CreateUser) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(params); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// TODO PENGECEKAN EMAIL SUDAH TERPAKAI ATAU BELUM
	user, _ := alc.userrepository.GetUsersByEmail(ctx, params.Email)
	if user.Email != "" {
		newError := errors.New("Email sudah digunakan")
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateUsers : %s", newError.Error()), newError)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  newError,
		}
	}
	hashPass, errHash := utils.HashPassword(params.KataSandi)
	if errHash != nil {
		log.Println(errHash)
		err = &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errHash,
		}
		return
	}
	// fmt.Println("[USECASE]", params.KataSandi, hashPass)
	tmpstp := utils.ParseDateToGoTime(params.TanggalLahir, err)
	tanggalLahir := utils.ParseDateToStr(tmpstp)
	resRepo, errRepo := alc.userrepository.CreateUsers(ctx, entity.User{
		Email:        params.Email,
		KataSandi:    hashPass,
		NamaUser:     params.NamaUser,
		NoTelp:       params.NoTelp,
		TanggalLahir: tanggalLahir,
		Pekerjaan:    params.Pekerjaan,
		IDProvinsi:   params.IDProvinsi,
		IDKota:       params.IDKota,
	})
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateUsers : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}