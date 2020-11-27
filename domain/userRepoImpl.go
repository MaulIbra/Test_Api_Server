package domain

import (
	"database/sql"
	"github.com/MaulIbra/Test_Api_Server/model"
	"github.com/MaulIbra/Test_Api_Server/utils"
	logs "github.com/MaulIbra/logs_module"
	guuid "github.com/google/uuid"
	"log"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &userRepo{
		db: db,
	}
}
func (u userRepo) CreateUser(user *model.User) error {
	id := guuid.New()
	user.IdUser = id.String()
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.CREATE_USER)
	defer stmt.Close()
	if err != nil {
		_ = tx.Rollback()
		logs.ErrorLogger.Println(err)
		return err
	}
	_, err = stmt.Exec(id, user.Nik, user.Username, user.TglLahir, user.Pekerjaan.IdPekerjaan, user.Pendidikan.IdPendidikan)
	if err != nil {
		_ = tx.Rollback()
		logs.ErrorLogger.Println(err)
		return err
	}

	return tx.Commit()
}

func (u userRepo) ReadUser(i int, i2 int) ([]*model.User, error) {
	users := make([]*model.User, 0)
	stmt, err := u.db.Prepare(utils.SELECT_USERS)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(i, i2)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(
			&user.IdUser, &user.Nik, &user.Username, &user.TglLahir, &user.Pendidikan.IdPendidikan, &user.Pendidikan.LabelPendidikan,
			&user.Pekerjaan.IdPekerjaan, &user.Pekerjaan.LabelPekerjaan, &user.UserStatus, &user.CreatedDate, &user.UpdatedDate)
		if err != nil {
			log.Print(err)
			logs.ErrorLogger.Println(err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u userRepo) ReadUserById(s string) (*model.User, error) {
	user := model.User{}
	stmt, err := u.db.Prepare(utils.SELECT_USER_BY_ID)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(s).Scan(
		&user.IdUser, &user.Nik, &user.Username, &user.TglLahir, &user.Pendidikan.IdPendidikan, &user.Pendidikan.LabelPendidikan,
		&user.Pekerjaan.IdPekerjaan, &user.Pekerjaan.LabelPekerjaan, &user.UserStatus, &user.CreatedDate, &user.UpdatedDate)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return &user, err
	}
	return &user, nil
}

func (u userRepo) UpdateUser(user *model.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(user.Nik, user.Username, user.TglLahir, user.Pekerjaan.IdPekerjaan, user.Pendidikan.IdPendidikan, user.IdUser)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u userRepo) DeleteUser(s string) error {
	tx, err := u.db.Begin()
	status := 0
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.DELETE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(status, s)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u userRepo) ReadPekerjaan() ([]*model.Pekerjaan, error) {
	listPekerjaan := make([]*model.Pekerjaan, 0)
	stmt, err := u.db.Prepare(utils.SELECT_PEKERJAAN)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}

	for rows.Next() {
		pekerjaan := model.Pekerjaan{}
		err := rows.Scan(&pekerjaan.IdPekerjaan, &pekerjaan.LabelPekerjaan)
		if err != nil {
			logs.ErrorLogger.Println(err)
			return nil, err
		}
		listPekerjaan = append(listPekerjaan, &pekerjaan)
	}
	return listPekerjaan, nil
}

func (u userRepo) ReadPendidikan() ([]*model.Pendidikan, error) {
	listPendidikan := make([]*model.Pendidikan, 0)
	stmt, err := u.db.Prepare(utils.SELECT_PENDIDIKAN)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		logs.ErrorLogger.Println(err)
		return nil, err
	}

	for rows.Next() {
		pendidikan := model.Pendidikan{}
		err := rows.Scan(&pendidikan.IdPendidikan, &pendidikan.LabelPendidikan)
		if err != nil {
			logs.ErrorLogger.Println(err)
			return nil, err
		}
		listPendidikan = append(listPendidikan, &pendidikan)
	}
	return listPendidikan, nil
}
