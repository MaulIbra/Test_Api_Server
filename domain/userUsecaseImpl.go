package domain

import (
	"github.com/MaulIbra/Test_Api_Server/model"
	logs "github.com/MaulIbra/logs_module"
)

type userUsecase struct {
	repo IUserRepo
}

func (u userUsecase) CreateUser(user *model.User) (*model.User, error) {
	err := u.repo.CreateUser(user)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	userResponse, err := u.repo.ReadUserById(user.IdUser)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return userResponse, nil
}

func (u userUsecase) ReadUser(i int, i2 int) ([]*model.User, error) {
	indexFirst := (i * i2) - i2
	indexLast := (i * i2) - 1
	users, err := u.repo.ReadUser(indexFirst, indexLast)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return users, nil
}

func (u userUsecase) ReadUserById(s string) (*model.User, error) {
	user,err := u.repo.ReadUserById(s)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return user,nil
}

func (u userUsecase) UpdateUser(user *model.User) error {
	err := u.repo.UpdateUser(user)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return err
	}
	return nil
}

func (u userUsecase) DeleteUser(s string) error {
	err := u.repo.DeleteUser(s)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return err
	}
	return nil
}

func (u userUsecase) ReadPekerjaan() ([]*model.Pekerjaan, error) {
	listPekerjaan,err := u.repo.ReadPekerjaan()
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return listPekerjaan,nil
}

func (u userUsecase) ReadPendidikan() ([]*model.Pendidikan, error) {
	listPendidikan,err := u.repo.ReadPendidikan()
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return listPendidikan,nil
}

func NewUserUsecase(repo IUserRepo) IUserUsecase {
	return &userUsecase{repo: repo}
}
