package auth

import (
	"database/sql"
	"github.com/MaulIbra/Test_Api_Server/model"
	"github.com/MaulIbra/Test_Api_Server/utils"
)

type authRepo struct {
	db *sql.DB
}

func (a authRepo) ReadAccountByEmail(s string) (*model.Account, error) {
	account := model.Account{}
	stmt, err := a.db.Prepare(utils.SELECT_ACCOUNT_BY_EMAIL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(s).Scan(&account.AccountId,&account.Email,&account.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return &account, nil
		}else{
			return &account, err
		}
	}
	return &account, nil
}

func (a authRepo) CreateAccount(account *model.Account,id string) error {

	tx, err := a.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.CREATE_ACCOUNT)

	if err != nil {
		_ = tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id,account.Email,account.Password)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	stmt.Close()
	return tx.Commit()
}

func NewAuthRepo(db *sql.DB) IAuthRepo {
	return &authRepo{db: db}
}