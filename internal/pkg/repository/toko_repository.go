package repository

import (
	"context"
	"errors"
	"fmt"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type TokoRepository interface {
	GetAllToko(ctx context.Context, params entity.FilterToko) (res []entity.Toko, err error)
	GetTokoByID(ctx context.Context, tokoId string) (res entity.Toko, err error)
	CreateToko(ctx context.Context, data entity.Toko) (res uint, err error)
	UpdateTokoByID(ctx context.Context, tokoId string, data entity.Toko) (res string, err error)
	DeleteTokoByID(ctx context.Context, tokoId string, userId string) (res string, err error)
	GetTokoByUserID(ctx context.Context, userID string) (res entity.Toko, err error)
}

type TokoRepositoryImpl struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &TokoRepositoryImpl{
		db: db,
	}
}
func (r *TokoRepositoryImpl) GetAllToko(ctx context.Context, params entity.FilterToko) (res []entity.Toko, err error) {
	db := r.db

	if params.Name != "" {
		db = db.Where("nama_toko like ?", "%"+params.Name)
	}

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *TokoRepositoryImpl) GetTokoByID(ctx context.Context, tokoId string) (res entity.Toko, err error) {
	if err := r.db.First(&res, tokoId).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *TokoRepositoryImpl) CreateToko(ctx context.Context, data entity.Toko) (res uint, err error) {
	result := r.db.Debug().Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (r *TokoRepositoryImpl) UpdateTokoByID(ctx context.Context, tokoId string, data entity.Toko) (res string, err error) {
	var dataToko entity.Toko
	if err = r.db.Where("id = ? ", tokoId).First(&dataToko).WithContext(ctx).Error; err != nil {
		return "Update books failed", gorm.ErrRecordNotFound
	}

	if err := r.db.Model(dataToko).Updates(&data).Where("id = ? ", tokoId).Error; err != nil {
		return "Update books failed", err
	}

	return res, nil
}

func (r *TokoRepositoryImpl) DeleteTokoByID(ctx context.Context, tokoId string, userId string) (res string, err error) {
	var dataToko entity.Toko
	if err = r.db.Where("id = ?", tokoId).First(&dataToko).WithContext(ctx).Error; err != nil {
		return "Delete book failed", gorm.ErrRecordNotFound
	}
	if fmt.Sprint(dataToko.IDUser) != userId {
		return "Delete book failed", errors.New("No Authorization")
	}

	if err := r.db.Model(dataToko).Delete(&dataToko).Error; err != nil {
		return "Delete book failed", err
	}

	return res, nil
}
func (r *TokoRepositoryImpl) GetTokoByUserID(ctx context.Context, userID string) (res entity.Toko, err error) {
	if err := r.db.Where("id_user = ?", userID).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}
