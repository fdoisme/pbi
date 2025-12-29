package repository

import (
	"context"
	"errors"
	"fmt"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/utils"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	GetAllProduk(ctx context.Context, params entity.FilterProduk) (res []entity.Produk, err error)
	GetProdukByID(ctx context.Context, produkId string) (res entity.Produk, err error)
	CreateProduk(ctx context.Context, data entity.Produk) (res uint, err error)
	UpdateProdukByID(ctx context.Context, produkId string, data entity.Produk) (res string, err error)
	DeleteProdukByID(ctx context.Context, produkId string, userId string) (res string, err error)
}

type ProdukRepositoryImpl struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &ProdukRepositoryImpl{
		db: db,
	}
}
func (r *ProdukRepositoryImpl) GetAllProduk(ctx context.Context, params entity.FilterProduk) (res []entity.Produk, err error) {
	db := r.db
	db.Debug().Preload("Toko").Preload("Category").Preload("Foto")
	if params.NamaProduk != "" {
		db.Where("nama_produk like ?", params.NamaProduk+"%")
	}
	if params.CategoryID != "" {
		x := utils.StrToUint(params.CategoryID)
		db.Where("id_category = ?", x)
	}
	if params.TokoID != "" {
		x := utils.StrToUint(params.TokoID)
		db.Where("id_toko = ?", x)
	}
	if params.MinHarga != "" {
		x := utils.StrToInt(params.MinHarga)
		db.Where("harga_reseller > ? OR harga_konsumen > ?", x, x)
	}
	if params.MaxHarga != "" {
		x := utils.StrToInt(params.MaxHarga)
		db.Where("harga_reseller < ? OR harga_konsumen < ?", x, x)
	}
	if err := db.WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *ProdukRepositoryImpl) GetProdukByID(ctx context.Context, produkId string) (res entity.Produk, err error) {
	if err := r.db.Preload("Toko").Preload("Category").Preload("Foto").First(&res, produkId).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *ProdukRepositoryImpl) CreateProduk(ctx context.Context, data entity.Produk) (res uint, err error) {
	result := r.db.Debug().Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (r *ProdukRepositoryImpl) UpdateProdukByID(ctx context.Context, produkId string, data entity.Produk) (res string, err error) {
	var dataProduk entity.Produk
	if err = r.db.Where("id = ? ", produkId).First(&dataProduk).WithContext(ctx).Error; err != nil {
		return "Update books failed", gorm.ErrRecordNotFound
	}

	if err := r.db.Model(dataProduk).Updates(&data).Where("id_produk = ? ", produkId).Error; err != nil {
		return "Update books failed", err
	}

	return res, nil
}

func (r *ProdukRepositoryImpl) DeleteProdukByID(ctx context.Context, produkId string, userId string) (res string, err error) {
	var dataProduk entity.Produk
	if err = r.db.Preload("Toko").Where("id = ?", produkId).First(&dataProduk).WithContext(ctx).Error; err != nil {
		return "Delete book failed", gorm.ErrRecordNotFound
	}
	if fmt.Sprint(dataProduk.Toko.IDUser) != userId {
		return "Delete book failed", errors.New("No Authorization")
	}
	if err := r.db.Model(dataProduk).Delete(&dataProduk).Error; err != nil {
		return "Delete book failed", err
	}

	return res, nil
}
