package user

import (
	"bytes"
	"encoding/json"
	"github.com/MaulIbra/Test_Api_Server/config"
	jwtToken "github.com/MaulIbra/go_module_jwtToken"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockRouteUser() *mux.Router {
	env := config.NewEnv()
	db := config.InitDB(env)
	userRepo := NewUserRepo(db)
	userUsecase := NewUserUsecase(userRepo)
	uc := UserController{userUsecase}
	router := mux.NewRouter()
	userRoute := router.PathPrefix("/user").Subrouter()
	userRoute.Use(jwtToken.TokenValidation)
	userRoute.HandleFunc("", uc.CreateUser).Methods(http.MethodPost)
	userRoute.HandleFunc("/{page}/{limit}", uc.ReadUser).Methods(http.MethodGet)
	userRoute.HandleFunc("/{id}", uc.ReadUserById).Methods(http.MethodGet)
	userRoute.HandleFunc("/{id}", uc.UpdateUser).Methods(http.MethodPut)
	userRoute.HandleFunc("/{id}", uc.DeleteUser).Methods(http.MethodDelete)

	etcRoute := router.PathPrefix("").Subrouter()
	etcRoute.Use(jwtToken.TokenValidation)
	etcRoute.HandleFunc("/job", uc.ReadJob).Methods(http.MethodGet)
	etcRoute.HandleFunc("/education", uc.ReadEducation).Methods(http.MethodGet)

	return router
}


func TestUserController_ReadUser(t *testing.T) {
	//mockup request body
	req, err := http.NewRequest("GET", "/user/1/3", nil)
	if err != nil {
		t.Fatalf("error %v", err)
	}

	response := httptest.NewRecorder()
	mockRouteUser().ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"Response is expected")
}

func TestUserController_CreateUser(t *testing.T) {
	//mockup request body
	byte,_ := json.Marshal(userMock)

	req, err := http.NewRequest("POST", "/user", bytes.NewReader(byte))
	if err != nil {
		t.Fatalf("error %v", err)
	}

	response := httptest.NewRecorder()
	mockRouteUser().ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"Response is expected")
}

func TestUserController_ReadUserById(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/36d0035f-cd53-4b55-a50c-0ce77d7b8ee4", nil)
	if err != nil {
		t.Fatalf("error %v", err)
	}

	response := httptest.NewRecorder()
	mockRouteUser().ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"Response is expected")
}

func TestUserController_UpdateUser(t *testing.T) {
	req, err := http.NewRequest("PUT", "/user/36d0035f-cd53-4b55-a50c-0ce77d7b8ee4", nil)
	if err != nil {
		t.Fatalf("error %v", err)
	}

	response := httptest.NewRecorder()
	mockRouteUser().ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"Response is expected")
}

func TestUserController_DeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user/36d0035f-cd53-4b55-a50c-0ce77d7b8ee4", nil)
	if err != nil {
		t.Fatalf("error %v", err)
	}

	response := httptest.NewRecorder()
	mockRouteUser().ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"Response is expected")
}

func TestUserController_ReadJob(t *testing.T) {
	req, err := http.NewRequest("GET", "/job", nil)
	if err != nil {
		t.Fatalf("error %v", err)
	}

	response := httptest.NewRecorder()
	mockRouteUser().ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"Response is expected")
}

func TestUserController_ReadEducation(t *testing.T) {
	req, err := http.NewRequest("GET", "/education", nil)
	if err != nil {
		t.Fatalf("error %v", err)
	}

	response := httptest.NewRecorder()
	mockRouteUser().ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"Response is expected")
}
