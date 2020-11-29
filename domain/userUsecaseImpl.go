package domain

import (
	"github.com/MaulIbra/Test_Api_Server/model"
	logs "github.com/MaulIbra/logs_module"
	"math"
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
	userResponse, err := u.repo.ReadUserById(user.UserId)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return userResponse, nil
}

func (u userUsecase) ReadUser(i int, i2 int) (*model.UserList, error) {
	indexFirst := (i * i2) - i2
	users, err := u.repo.ReadUser(indexFirst, i2)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}

	totalData, err := u.repo.CountUser()
	if err != nil {
		logs.ErrorLogger.Println(err)
	}

	totalPage := math.Ceil(float64(totalData)/float64(i2))

	return &model.UserList{
		Users: users,
		Metadata: model.Metadata{
			CurrentPage: i,
			FirstPage:   1,
			LastPage:    int(totalPage),
			TotalData:   totalData,
		},
	}, nil
}

func (u userUsecase) ReadUserById(s string) (*model.User, error) {
	user, err := u.repo.ReadUserById(s)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return user, nil
}

func (u userUsecase) UpdateUser(user *model.User) (*model.User, error) {
	err := u.repo.UpdateUser(user)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	userResp, err := u.repo.ReadUserById(user.UserId)
	if err != nil {
		return nil, err
	}
	return userResp, err
}

func (u userUsecase) DeleteUser(s string) error {
	err := u.repo.DeleteUser(s)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return err
	}
	return nil
}

func (u userUsecase) ReadJob() ([]*model.Job, error) {
	jobList, err := u.repo.ReadJob()
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return jobList, nil
}

func (u userUsecase) ReadEducation() ([]*model.Education, error) {
	educationList, err := u.repo.ReadEducation()
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	return educationList, nil
}

func NewUserUsecase(repo IUserRepo) IUserUsecase {
	return &userUsecase{repo: repo}
}
