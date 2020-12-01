package domain

import "github.com/MaulIbra/Test_Api_Server/model"

type IAuthUsecase interface {
	Login(account *model.Account)(*model.Account,error)
	Register(account *model.Account)(*model.Account,error)
}
