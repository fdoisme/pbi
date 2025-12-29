package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type FotoProdukRepository interface {
	CreateProduk(ctx context.Context, data entity.FotoProduk) (res uint, err error)
	UpdateProdukByProdukID(ctx context.Context, produkId string, data []string) (res string, err error)
}

type FotoProdukRepositoryImpl struct {
	db *gorm.DB
}

func NewFotoProdukRepository(db *gorm.DB) FotoProdukRepository {
	return &FotoProdukRepositoryImpl{
		db: db,
	}
}

func (r *FotoProdukRepositoryImpl) CreateProduk(ctx context.Context, data entity.FotoProduk) (res uint, err error) {
	result := r.db.Debug().Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}
	return data.ID, nil
}

func (r *FotoProdukRepositoryImpl) UpdateProdukByProdukID(ctx context.Context, produkId string, data []string) (res string, err error) {
	var dataFotoProduk []entity.FotoProduk
	if err = r.db.Where("id_produk = ? ", produkId).Find(&dataFotoProduk).WithContext(ctx).Error; err != nil {
		return "Update foto produk failed", gorm.ErrRecordNotFound
	}

	for i := range dataFotoProduk {
		if err := r.db.Model(&entity.FotoProduk{}).Where("id_produk = ? ", produkId).Update("url", data[i]).Error; err != nil {
			return "Update foto produk failed", err
		}
	}

	return res, nil
}
