package domain

import (
	"context"
	"github.com/MaulIbra/Test_Api_Server/config"
	"github.com/gorilla/mux"
)

type AppRoute struct {
	Router  *mux.Router
	Ctx     context.Context
}

func Init()  {
	route := AppRoute{}
	route.Router = config.MuxRouter()
	route.Ctx = context.Background()
	route.InitRoute()
	route.Serve()
}

func (route *AppRoute) Serve()  {
	config.ListenServe(route.Router)
}

func (route *AppRoute) InitRoute()  {
	env := config.NewEnv()
	db := config.InitDB(env)
	userRepo := NewUserRepo(db)
	userUsecase := NewUserUsecase(userRepo)
	userRoute := UserController{Route: route,Usecase: userUsecase}
	userRoute.UserRoute()
}

