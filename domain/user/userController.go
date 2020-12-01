package user

import (
	"github.com/MaulIbra/Test_Api_Server/model"
	"github.com/MaulIbra/Test_Api_Server/utils"
	jwtToken "github.com/MaulIbra/go_module_jwtToken"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserController struct {
	Usecase IUserUsecase
}

func (uc *UserController) UserRoute(router *mux.Router) {
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
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := utils.JsonDecoder(&user, r)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	isValid := user.Validate()
	if !isValid{
		utils.Response(w, http.StatusBadRequest, user)
		return
	}
	user.UserId = guuid.New().String()
	userResponse, err := uc.Usecase.CreateUser(&user)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	utils.Response(w, http.StatusCreated, userResponse)
}

func (uc *UserController) ReadUser(w http.ResponseWriter, r *http.Request) {
	if len(utils.DecodePathVariabel("page", r)) == 0 {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	page, _ := strconv.Atoi(utils.DecodePathVariabel("page", r))

	if len(utils.DecodePathVariabel("page", r)) == 0 {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	limit, _ := strconv.Atoi(utils.DecodePathVariabel("limit", r))

	users, err := uc.Usecase.ReadUser(page, limit)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	utils.Response(w, http.StatusOK, users)
}

func (uc *UserController) ReadUserById(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	if len(id) == 0 {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	user, err := uc.Usecase.ReadUserById(id)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	if user.Job == (model.Job{}){
		utils.Response(w, http.StatusNoContent,user)
		return
	}
	utils.Response(w, http.StatusOK, user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	id := utils.DecodePathVariabel("id", r)
	if len(id) == 0 {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	user.UserId = id
	err := utils.JsonDecoder(&user, r)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	isValid := user.Validate()
	if !isValid{
		utils.Response(w, http.StatusBadRequest, user)
		return
	}
	userResp, err := uc.Usecase.UpdateUser(&user)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	if userResp.Job == (model.Job{}){
		utils.Response(w, http.StatusNoContent,userResp)
		return
	}
	utils.Response(w, http.StatusOK, userResp)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	if len(id) == 0 {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	err := uc.Usecase.DeleteUser(id)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	utils.ResponseWithoutPayload(w, http.StatusOK)
}

func (uc *UserController) ReadJob(w http.ResponseWriter, r *http.Request) {
	job, err := uc.Usecase.ReadJob()
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	utils.Response(w, http.StatusOK, job)
}

func (uc *UserController) ReadEducation(w http.ResponseWriter, r *http.Request) {
	education, err := uc.Usecase.ReadEducation()
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	utils.Response(w, http.StatusOK, education)
}
