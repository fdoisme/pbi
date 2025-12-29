package mapper

import (
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
)

func MapperToProdukRespon(data entity.Produk) model.ProdukRes {
	var result model.ProdukRes
	if len(data.Foto) > 0  {
		for _, v := range data.Foto {
			photosRes := model.FotoProdukRes{
				ID: v.ID,
				IDProduk: v.IDProduk,
				URL: v.URL,
			}
			result.Photos = append(result.Photos, photosRes)
		}
	}
	result.Category.ID = data.Category.ID
	result.Category.NamaCategory = data.Category.NamaCategory
	result.Toko.ID = data.Toko.ID
	result.Toko.NamaToko = data.Toko.NamaToko
	result.Toko.URLFoto = data.Toko.URLFoto
	result.ID = data.ID
	result.NamaProduk = data.NamaProduk
	result.Deskripsi = data.Deskripsi
	result.Slug = data.Slug
	result.HargaKonsumen = data.HargaKonsumen
	result.HargaReseller = data.HargaReseller
	result.Stok = data.Stok
	return result
}
