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
	user.UserId = id.String()
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
	_, err = stmt.Exec(id, user.IdCardNumber, user.Username, user.DateOfBirth, user.Job.JobId, user.Education.EducationId)
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
			&user.UserId, &user.IdCardNumber, &user.Username, &user.DateOfBirth, &user.Education.EducationId, &user.Education.EducationLabel,
			&user.Job.JobId, &user.Job.JobLabel, &user.UserStatus, &user.CreatedDate, &user.UpdatedDate)
		if err != nil {
			log.Print(err)
			logs.ErrorLogger.Println(err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u userRepo) CountUser() (int, error) {
	var totalData int
	stmt, err := u.db.Prepare(utils.SELECT_COUNT_DATA_USER)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return totalData, err
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&totalData)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return totalData, err
	}
	return totalData, nil
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
		&user.UserId, &user.IdCardNumber, &user.Username, &user.DateOfBirth, &user.Education.EducationId, &user.Education.EducationLabel,
		&user.Job.JobId, &user.Job.JobLabel, &user.UserStatus, &user.CreatedDate, &user.UpdatedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, nil
		}else{
			logs.ErrorLogger.Println(err)
			return &user, err
		}
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
	_, err = stmt.Exec(user.IdCardNumber, user.Username, user.DateOfBirth, user.Job.JobId, user.Education.EducationId, user.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u userRepo) DeleteUser(s string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.DELETE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(s)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u userRepo) ReadJob() ([]*model.Job, error) {
	jobList := make([]*model.Job, 0)
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
		job := model.Job{}
		err := rows.Scan(&job.JobId, &job.JobLabel)
		if err != nil {
			logs.ErrorLogger.Println(err)
			return nil, err
		}
		jobList = append(jobList, &job)
	}
	return jobList, nil
}

func (u userRepo) ReadEducation() ([]*model.Education, error) {
	educationList := make([]*model.Education, 0)
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
		education := model.Education{}
		err := rows.Scan(&education.EducationId, &education.EducationLabel)
		if err != nil {
			logs.ErrorLogger.Println(err)
			return nil, err
		}
		educationList = append(educationList, &education)
	}
	return educationList, nil
}
