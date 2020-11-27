package model

type User struct {
	idUser     string       `json:"idUser"`
	Nik        string       `json:"nik"`
	TglLahir   TanggalLahir `json:"tglLahir"`
	Pekerjaan  Pekerjaan    `json:"pekerjaan"`
	Pendidikan Pendidikan   `json:"pendidikan"`
}

type Pekerjaan struct {
	IdPekerjaan    string `json:"idPekerjaan"`
	LabelPekerjaan string `json:"labelPekerjaan"`
}

type Pendidikan struct {
	IdPendidikan    string `json:"idPendidikan"`
	LabelPendidikan string `json:"labelPendidikan"`
}

type TanggalLahir struct {
	Tanggal int `json:"tanggal"`
	Bulan   int `json:"bulan"`
	Tahun   int `json:"tahun"`
}
