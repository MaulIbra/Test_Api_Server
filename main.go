package main

import (
	"github.com/MaulIbra/Test_Api_Server/domain"
	logs "github.com/MaulIbra/logs_module"
)

func main() {
	logs.InitLog("log")
	domain.Init()
}
