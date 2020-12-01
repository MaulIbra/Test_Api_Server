package domain

import (
	"errors"
	"github.com/MaulIbra/Test_Api_Server/model"
	encrypt "github.com/MaulIbra/encrypt_module"
	jwtToken "github.com/MaulIbra/go_module_jwtToken"
	"sync"
)

type authUsecase struct {
	authRepo IAuthRepo
}

func (a authUsecase) Login(account *model.Account) (*model.Account, error) {
	token := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go jwtToken.GenerateToken(7200, token, &wg)

	accountResp, err := a.authRepo.ReadAccountByEmail(account.Email)
	if err != nil {
		return nil, err
	}
	if accountResp.Email == "" {
		return &model.Account{}, errors.New("Account not yet registered")
	}

	go func() {
		wg.Wait()
		close(token)
	}()

	authenticate := encrypt.CompareEncrypt(accountResp.Password, []byte(account.Password))
	if authenticate {
		return &model.Account{
			Email: account.Email,
			Token: <-token,
		}, nil
	}
	return &model.Account{}, errors.New("Email Or Password was invalid")
}

func (a authUsecase) Register(account *model.Account) (*model.Account, error) {
	token := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go jwtToken.GenerateToken(7200, token, &wg)

	encryptPass := encrypt.Encrypt([]byte(account.Password))
	account.Password = encryptPass

	err := a.authRepo.CreateAccount(account)
	if err != nil {
		return nil, err
	}
	go func() {
		wg.Wait()
		close(token)
	}()
	return &model.Account{
		AccountId: account.AccountId,
		Email:     account.Email,
		Token:     <-token,
	}, nil
}

func NewAuthUsecase(authRepo IAuthRepo) IAuthUsecase {
	return &authUsecase{authRepo: authRepo}
}
