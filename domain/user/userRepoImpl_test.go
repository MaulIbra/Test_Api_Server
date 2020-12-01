package user

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

var userMock = &model.User{
	UserId:       uuid.New().String(),
	IdCardNumber: "3209120503980004",
	Username:     "Maulana Ibrahim",
	DateOfBirth:  "1998-03-05",
	Job: model.Job{
		JobId:    uuid.New().String(),
		JobLabel: "Wiraswasta",
	},
	Education: model.Education{
		EducationId:    uuid.New().String(),
		EducationLabel: "Diploma",
	},
	UserStatus:  1,
	CreatedDate: "2020-12-01 09:36:44",
	UpdatedDate: "2020-12-01 09:36:44",
}

func TestCreateUser(t *testing.T) {
	db, mock := config.NewMock()
	repo := userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.CREATE_USER)).ExpectExec().WithArgs(userMock.UserId, userMock.IdCardNumber, userMock.Username,userMock.DateOfBirth,userMock.Job.JobId,userMock.Education.EducationId).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.CreateUser(userMock)
	assert.NoError(t, err)
}

func TestCreateUserError(t *testing.T) {
	db, mock := config.NewMock()
	repo := userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.CREATE_USER)).ExpectExec().WithArgs(userMock.UserId, userMock.IdCardNumber, userMock.Username,userMock.DateOfBirth,userMock.Job.JobId,userMock.Education.EducationId).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.CreateUser(userMock)
	assert.Error(t, err)
}

func TestReadUser(t *testing.T){
	db, mock := config.NewMock()
	repo := &userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id_user", "nik", "username","tgl_lahir","id_pendidikan","label_pendidikan","id_pekerjaan","label_pekerjaan","user_status","created_date","update_date"}).
		AddRow(userMock.UserId,userMock.IdCardNumber,userMock.Username,userMock.DateOfBirth,userMock.Education.EducationId,userMock.Education.EducationLabel,userMock.Job.JobId,userMock.Job.JobLabel,userMock.UserStatus,userMock.CreatedDate,userMock.UpdatedDate)

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_USERS)).ExpectQuery().WithArgs(1,3).WillReturnRows(rows)

	user, err := repo.ReadUser(1,3)
	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.Len(t,user,1)
}

func TestCountUser(t *testing.T){
	db,mock := config.NewMock()
	repo := &userRepo{db: db}

	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"total_data"}).
		AddRow(2)

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_COUNT_DATA_USER)).ExpectQuery().WillReturnRows(rows)
	totalData,err := repo.CountUser()
	assert.NotNil(t, totalData)
	assert.NoError(t, err)
	assert.Equal(t,2,totalData)
}

func TestCountUserError(t *testing.T){
	db,mock := config.NewMock()
	repo := &userRepo{db: db}

	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"total_data"})

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_COUNT_DATA_USER)).ExpectQuery().WillReturnRows(rows)
	totalData,err := repo.CountUser()
	assert.Error(t, err)
	assert.NotEqual(t,2,totalData)
}

func TestReadUserById(t *testing.T){
	db, mock := config.NewMock()
	repo := &userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id_user", "nik", "username","tgl_lahir","id_pendidikan","label_pendidikan","id_pekerjaan","label_pekerjaan","user_status","created_date","update_date"}).
		AddRow(userMock.UserId,userMock.IdCardNumber,userMock.Username,userMock.DateOfBirth,userMock.Education.EducationId,userMock.Education.EducationLabel,userMock.Job.JobId,userMock.Job.JobLabel,userMock.UserStatus,userMock.CreatedDate,userMock.UpdatedDate)

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_USER_BY_ID)).ExpectQuery().WithArgs(userMock.UserId).WillReturnRows(rows)

	user, err := repo.ReadUserById(userMock.UserId)
	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.NotEmpty(t,user)
}

func TestReadUserByIdError(t *testing.T){
	db, mock := config.NewMock()
	repo := &userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id_user", "nik", "username","tgl_lahir","id_pendidikan","label_pendidikan","id_pekerjaan","label_pekerjaan","user_status","created_date","update_date"})
	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_USER_BY_ID)).ExpectQuery().WithArgs(userMock.UserId).WillReturnRows(rows)

	user, err := repo.ReadUserById(userMock.UserId)
	assert.NoError(t, err)
	assert.Nil(t,err)
	assert.Empty(t, user)
}

func TestUpdateUser(t *testing.T) {
	db, mock := config.NewMock()
	repo := userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.UPDATE_USER)).ExpectExec().WithArgs(userMock.IdCardNumber, userMock.Username,userMock.DateOfBirth,userMock.Job.JobId,userMock.Education.EducationId,userMock.UserId).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.UpdateUser(userMock)
	assert.NoError(t, err)
}

func TestUpdateUserError(t *testing.T) {
	db, mock := config.NewMock()
	repo := userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.UPDATE_USER)).ExpectExec().WithArgs(userMock.IdCardNumber, userMock.Username,userMock.DateOfBirth,userMock.Job.JobId,userMock.Education.EducationId,userMock.UserId).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.UpdateUser(userMock)
	assert.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, mock := config.NewMock()
	repo := userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.DELETE_USER)).ExpectExec().WithArgs(userMock.UserId).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.DeleteUser(userMock.UserId)
	assert.NoError(t, err)
}

func TestDeleteUserError(t *testing.T) {
	db, mock := config.NewMock()
	repo := userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.DELETE_USER)).ExpectExec().WithArgs(userMock.UserId).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectRollback()

	err := repo.DeleteUser(userMock.UserId)
	assert.Error(t, err)
}

func TestReadJob(t *testing.T){
	db, mock := config.NewMock()
	repo := &userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id_pekerjaan", "label_pekerjaan"}).
		AddRow(userMock.Job.JobId,userMock.Job.JobLabel)

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_PEKERJAAN)).ExpectQuery().WithArgs().WillReturnRows(rows)

	job, err := repo.ReadJob()
	assert.NotNil(t, job)
	assert.NoError(t, err)
	assert.Len(t,job,1)
}

func TestReadJobError(t *testing.T){
	db, mock := config.NewMock()
	repo := &userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id_pekerjaan", "label_pekerjaan"}).
		AddRow(userMock.Job.JobId,userMock.Job.JobLabel)

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_PEKERJAAN)).ExpectQuery().WithArgs(1,3).WillReturnRows(rows)

	job, err := repo.ReadJob()
	assert.Nil(t, job)
	assert.Error(t, err)
	assert.Len(t,job,0)
}

func TestReadEducation(t *testing.T){
	db, mock := config.NewMock()
	repo := &userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id_pendidikan", "label_pendidikan"}).
		AddRow(userMock.Job.JobId,userMock.Job.JobLabel)

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_PENDIDIKAN)).ExpectQuery().WithArgs().WillReturnRows(rows)

	job, err := repo.ReadEducation()
	assert.NotNil(t, job)
	assert.NoError(t, err)
	assert.Len(t,job,1)
}

func TestReadEducationError(t *testing.T){
	db, mock := config.NewMock()
	repo := &userRepo{db}
	defer func() {
		repo.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id_pendidikan", "label_pendidikan"}).
		AddRow(userMock.Job.JobId,userMock.Job.JobLabel)

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_PENDIDIKAN)).ExpectQuery().WithArgs(1,3).WillReturnRows(rows)

	job, err := repo.ReadEducation()
	assert.Nil(t, job)
	assert.Error(t, err)
	assert.Len(t,job,0)
}