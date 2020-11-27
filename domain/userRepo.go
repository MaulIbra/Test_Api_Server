package domain

import "github.com/MaulIbra/Test_Api_Server/model"

type IUserRepo interface {
	CreateUser(*model.User) error
	ReadUser(int, int) ([]*model.User, error)
	ReadUserById(string) (*model.User, error)
	UpdateUser(*model.User) error
	DeleteUser(string) error
	ReadPekerjaan() ([]*model.Pekerjaan, error)
	ReadPendidikan() ([]*model.Pendidikan, error)
}
