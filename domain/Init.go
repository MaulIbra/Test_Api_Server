package domain

import (
	"context"
	"github.com/MaulIbra/Test_Api_Server/config"
	"github.com/MaulIbra/Test_Api_Server/domain/auth"
	"github.com/MaulIbra/Test_Api_Server/domain/user"
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

	userRepo := user.NewUserRepo(db)
	userUsecase := user.NewUserUsecase(userRepo)
	userRoute := user.UserController{Usecase: userUsecase}
	userRoute.UserRoute(route.Router)

	authRepo := auth.NewAuthRepo(db)
	authUsecase := auth.NewAuthUsecase(authRepo)
	authRoute := auth.AuthController{Usecase: authUsecase}
	authRoute.Auth(route.Router)
}

