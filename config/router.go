package config

import (
	"fmt"
	"github.com/MaulIbra/Test_Api_Server/utils"
	logs "github.com/MaulIbra/logs_module"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func MuxRouter() *mux.Router {
	return mux.NewRouter()
}

func ListenServe(router *mux.Router) {
	port := os.Getenv("SERVER_PORT")
	if os.Getenv("SERVER_PORT") == "" {
		port = utils.GetEnv("SERVER_PORT", "4000")
	}

	host := os.Getenv("SERVER_HOST")
	if os.Getenv("SERVER_HOST") == "" {
		port = utils.GetEnv("SERVER_HOST", "localhost")
	}

	log.Print(fmt.Sprintf("%v:%v", host, port))
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), router)
	if err != nil {
		panic(err)
		logs.ErrorLogger.Println(err)
	}
}
