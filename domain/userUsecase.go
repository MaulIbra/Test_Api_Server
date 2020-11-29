package domain

import "github.com/MaulIbra/Test_Api_Server/model"

type IUserUsecase interface {
	CreateUser(*model.User) (*model.User,error)
	ReadUser(int, int) (*model.UserList, error)
	ReadUserById(string) (*model.User, error)
	UpdateUser(*model.User) (*model.User,error)
	DeleteUser(string) error
	ReadJob() ([]*model.Job, error)
	ReadEducation() ([]*model.Education, error)
}
