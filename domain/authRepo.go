package domain

import "github.com/MaulIbra/Test_Api_Server/model"

type IAuthRepo interface {
	ReadAccountByEmail(string)(*model.Account,error)
	CreateAccount(account *model.Account) error
}
