package model

type User struct {
	IdUser      string     `json:"idUser"`
	Nik         string     `json:"nik"`
	Username    string     `json:"username"`
	TglLahir    string     `json:"tglLahir"`
	Pekerjaan   Pekerjaan  `json:"pekerjaan"`
	Pendidikan  Pendidikan `json:"pendidikan"`
	UserStatus  int        `json:"userStatus"`
	CreatedDate string     `json:"createdDate"`
	UpdatedDate string     `json:"updatedDate"`
}

type Pekerjaan struct {
	IdPekerjaan    string `json:"idPekerjaan"`
	LabelPekerjaan string `json:"labelPekerjaan"`
}

type Pendidikan struct {
	IdPendidikan    string `json:"idPendidikan"`
	LabelPendidikan string `json:"labelPendidikan"`
}
