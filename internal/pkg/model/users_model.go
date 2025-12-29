package model

type CreateUser struct {
	Email        string `json:"email" validate:"required,email"`
	KataSandi    string `json:"kata_sandi" validate:"required"`
	NamaUser     string `json:"nama" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required,min=10,max=13"`
	TanggalLahir string `json:"tanggal_Lahir" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	IDProvinsi   string `json:"id_provinsi" validate:"required,number"`
	IDKota       string `json:"id_kota" validate:"required,number"`
}

type Login struct {
	NoTelp   string `json:"no_telp" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type LoginRes struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
