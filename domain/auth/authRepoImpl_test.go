package auth

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MaulIbra/Test_Api_Server/config"
	"github.com/MaulIbra/Test_Api_Server/model"
	"github.com/MaulIbra/Test_Api_Server/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var account = &model.Account{
	AccountId: uuid.New().String(),
	Email:     "maulibrahim19@gmail.com",
	Password:  "maulana",
}

func TestReadAccountByEmail(t *testing.T) {
	db, mock := config.NewMock()
	repo := &authRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"account_id", "email", "password"}).
		AddRow(account.AccountId, account.Email, account.Password)

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_ACCOUNT_BY_EMAIL)).ExpectQuery().WithArgs(account.Email).WillReturnRows(rows)

	account, err := repo.ReadAccountByEmail(account.Email)
	assert.NotNil(t, account)
	assert.NoError(t, err)
}
func TestReadAccountByEmailError(t *testing.T) {
	db, mock := config.NewMock()
	repo := &authRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"account_id", "email", "password"})

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_ACCOUNT_BY_EMAIL)).ExpectQuery().WithArgs(account.Email).WillReturnRows(rows)

	account, err := repo.ReadAccountByEmail(account.Email)
	assert.Empty(t, account)
	assert.NoError(t, err)
}

func TestCreateAccount(t *testing.T) {
	db, mock := config.NewMock()
	repo := &authRepo{db}
	defer func() {
		repo.db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.CREATE_ACCOUNT)).ExpectExec().WithArgs(account.AccountId, account.Email, account.Password).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.CreateAccount(account,account.AccountId)
	assert.NoError(t, err)
}


func TestCreateAccountError(t *testing.T) {
	db, mock := config.NewMock()
	repo := &authRepo{db}
	defer func() {
		repo.db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.CREATE_ACCOUNT)).ExpectExec().WithArgs(account.AccountId, account.Email, account.Password).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.CreateAccount(account,account.AccountId)
	assert.Error(t, err)
}
