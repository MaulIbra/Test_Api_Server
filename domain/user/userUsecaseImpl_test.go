package user

import (
	"github.com/MaulIbra/Test_Api_Server/config"
	logs "github.com/MaulIbra/logs_module"
	"github.com/stretchr/testify/assert"
	"testing"
)
func mockUsecaseUser() IUserUsecase {
	env := config.NewEnv()
	db := config.InitDB(env)
	userRepo := NewUserRepo(db)
	userUsecase := NewUserUsecase(userRepo)
	return userUsecase
}

func TestUserUsecase_DeleteUserError(t *testing.T) {
	logs.InitLog("user_test_mock")
	err := mockUsecaseUser().DeleteUser("676687687")
	assert.Nil(t,err)
}

//func TestUserUsecase_CreateUser(t *testing.T) {
//	logs.InitLog("user_test_mock")
//	userMock.IdCardNumber = "3209120503980054"
//	user,err := mockUsecaseUser().CreateUser(userMock)
//	assert.NoError(t,err)
//	assert.NotEmpty(t,user)
//}

func TestUserUsecase_CreateUserError(t *testing.T) {
	logs.InitLog("user_test_mock")
	user,err := mockUsecaseUser().CreateUser(userMock)
	assert.Error(t,err)
	assert.Nil(t,user)
}

func TestUserUsecase_ReadUser(t *testing.T) {
	logs.InitLog("user_test_mock")
	users,err := mockUsecaseUser().ReadUser(1,3)
	assert.NoError(t,err)
	assert.NotEmpty(t,users)
}

func TestUserUsecase_ReadUserError(t *testing.T) {
	logs.InitLog("user_test_mock")
	_,err := mockUsecaseUser().ReadUser(-1,-100)
	assert.Nil(t,err)
}

func TestUserUsecase_ReadUserById(t *testing.T) {
	logs.InitLog("user_test_mock")
	userId := "85c2f7ca-ae5e-481f-bfe6-d147f1b92ef4"
	user,err := mockUsecaseUser().ReadUserById(userId)
	assert.NoError(t,err)
	assert.Equal(t,userId,user.UserId)
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	logs.InitLog("user_test_mock")
	mockUsername := "mock update"
	userMock.Username = mockUsername
	users,err := mockUsecaseUser().UpdateUser(userMock)
	assert.NotNil(t,users)
	assert.NoError(t,err)
}


