package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type CategoriesRepository interface {
	GetAllCategories(ctx context.Context) (res []entity.Category, err error)
	GetCategoriesByID(ctx context.Context, categoryId string) (res entity.Category, err error)
	CreateCategories(ctx context.Context, data entity.Category) (res uint, err error)
	UpdateCategoriesByID(ctx context.Context, categoryId string, data entity.Category) (res string, err error)
	DeleteCategoriesByID(ctx context.Context, categoryId string) (res string, err error)
}

type CategoriesRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) CategoriesRepository {
	return &CategoriesRepositoryImpl{
		db: db,
	}
}
func (r *CategoriesRepositoryImpl) GetAllCategories(ctx context.Context) (res []entity.Category, err error) {
	db := r.db

	if err := db.Debug().WithContext(ctx).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *CategoriesRepositoryImpl) GetCategoriesByID(ctx context.Context, categoryId string) (res entity.Category, err error) {
	if err := r.db.First(&res, categoryId).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *CategoriesRepositoryImpl) CreateCategories(ctx context.Context, data entity.Category) (res uint, err error) {
	result := r.db.Debug().Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (r *CategoriesRepositoryImpl) UpdateCategoriesByID(ctx context.Context, categoryId string, data entity.Category) (res string, err error) {
	var dataCategories entity.Category
	if err = r.db.Where("id = ? ", categoryId).First(&dataCategories).WithContext(ctx).Error; err != nil {
		return "Update categories failed", gorm.ErrRecordNotFound
	}

	if err := r.db.Model(dataCategories).Updates(&data).Where("id = ? ", categoryId).Error; err != nil {
		return "Update categories failed", err
	}

	return res, nil
}

func (r *CategoriesRepositoryImpl) DeleteCategoriesByID(ctx context.Context, categoryId string) (res string, err error) {
	var dataCategories entity.Category
	if err = r.db.Where("id = ?", categoryId).First(&dataCategories).WithContext(ctx).Error; err != nil {
		return "record not found", gorm.ErrRecordNotFound
	}

	if err := r.db.Model(dataCategories).Delete(&dataCategories).Error; err != nil {
		return "Delete category failed", err
	}

	return res, nil
}
