package domain

import (
	"github.com/MaulIbra/Test_Api_Server/model"
	"github.com/MaulIbra/Test_Api_Server/utils"
	logs "github.com/MaulIbra/logs_module"
	"net/http"
)

type AuthController struct {
	Route   *AppRoute
	Usecase IAuthUsecase
}

func (ac *AuthController) Auth()  {
	account := ac.Route.Router.PathPrefix("/account").Subrouter()
	account.HandleFunc("/login",ac.Login).Methods(http.MethodPost)
	account.HandleFunc("/register",ac.Register).Methods(http.MethodPost)
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request){
	account := model.Account{}
	err := utils.JsonDecoder(&account, r)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	isValid := account.Validate()
	if !isValid{
		utils.Response(w, http.StatusBadRequest, account)
		return
	}
	accountResp, err := ac.Usecase.Login(&account)
	if (accountResp.Token == "") && (err != nil) {
		logs.ErrorLogger.Println(err)
		utils.ResponseCustom(w, http.StatusUnauthorized,err.Error(), accountResp)
		return
	} else if err != nil {
		logs.ErrorLogger.Println(err)
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}

	utils.Response(w, http.StatusOK, accountResp)
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request){
	account := model.Account{}
	err := utils.JsonDecoder(&account, r)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	isValid := account.Validate()
	if !isValid{
		utils.Response(w, http.StatusBadRequest, account)
		return
	}
	accountResp, err := ac.Usecase.Register(&account)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	utils.Response(w, http.StatusCreated, accountResp)
}
