package domain

import (
	"database/sql"
	"github.com/MaulIbra/Test_Api_Server/model"
	"github.com/MaulIbra/Test_Api_Server/utils"
	logs "github.com/MaulIbra/logs_module"
	guuid "github.com/google/uuid"
)

type authRepo struct {
	db *sql.DB
}

func (a authRepo) ReadAccountByEmail(s string) (*model.Account, error) {
	account := model.Account{}
	stmt, err := a.db.Prepare(utils.SELECT_ACCOUNT_BY_EMAIL)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(s).Scan(&account.AccountId,&account.Email,&account.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return &account, nil
		}else{
			logs.ErrorLogger.Println(err)
			return &account, err
		}
	}
	return &account, nil
}

func (a authRepo) CreateAccount(account *model.Account) error {
	id := guuid.New()
	account.AccountId = id.String()
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.CREATE_ACCOUNT)
	defer stmt.Close()
	if err != nil {
		_ = tx.Rollback()
		logs.ErrorLogger.Println(err)
		return err
	}
	_, err = stmt.Exec(id,account.Email,account.Password)
	if err != nil {
		_ = tx.Rollback()
		logs.ErrorLogger.Println(err)
		return err
	}
	return tx.Commit()
}

func NewAuthRepo(db *sql.DB) IAuthRepo  {
	return &authRepo{db: db}
}