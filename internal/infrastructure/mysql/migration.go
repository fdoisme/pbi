package mysql

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
        &entity.User{},
        &entity.Alamat{},
        &entity.Category{},
        &entity.Produk{},
        &entity.FotoProduk{},
        &entity.LogProduk{},
        &entity.Toko{},
        &entity.Transaction{},
        &entity.DetailTransaction{},
	)
	if err != nil {
		helper.Logger(helper.LoggerLevelError, "Failed Database Migrated", err)
	}
    // tables := []interface{}{
    //     &entity.User{},
    //     &entity.Alamat{},
    //     &entity.Category{},
    //     &entity.Produk{},
    //     &entity.FotoProduk{},
    //     &entity.LogProduk{},
    //     &entity.Toko{},
    //     &entity.Transaction{},
    //     &entity.DetailTransaction{},
    // }
    // for _, table := range tables {
    //     if !mysqlDB.Migrator().HasTable(table) {
    //         err := mysqlDB.AutoMigrate(table)
    //         if err != nil {
    //             helper.Logger(helper.LoggerLevelError, "Failed to migrate table", err)
    //             return
    //         }
    //     }
    // }
	var count int64
	if mysqlDB.Migrator().HasTable(&entity.Category{}) {
		mysqlDB.Model(&entity.Category{}).Count(&count)
		if count < 1 {
			mysqlDB.CreateInBatches(categoriesSeed, len(categoriesSeed))
		}
	}

	helper.Logger(helper.LoggerLevelInfo, "Database Migrated", nil)
}
