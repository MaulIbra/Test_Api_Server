package config

import (
	"github.com/MaulIbra/Test_Api_Server/utils"
	"os"
)

type Env struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	SchemaName string
	Driver     string
}

func NewEnv() *Env {
	env := Env{}
	env.DbUser = os.Getenv("DB_USER")
	if os.Getenv("DB_USER") == "" {
		env.DbUser = utils.GetEnv("DB_USER", "root")
	}

	env.DbPassword = os.Getenv("DB_PASSWORD")
	if os.Getenv("DB_PASSWORD") == "" {
		env.DbPassword = utils.GetEnv("DB_PASSWORD", "")
	}

	env.DbHost = os.Getenv("DB_HOST")
	if os.Getenv("DB_HOST") == "" {
		env.DbHost = utils.GetEnv("DB_HOST", "127.0.0.1")
	}

	env.DbPort = os.Getenv("DB_PORT")
	if os.Getenv("DB_PORT") == "" {
		env.DbPort = utils.GetEnv("DB_PORT", "3306")
	}

	env.SchemaName = os.Getenv("DB_SCHEMA")
	if os.Getenv("DB_SCHEMA") == "" {
		env.SchemaName = utils.GetEnv("DB_SCHEMA", "db_user")
	}

	env.Driver = os.Getenv("DB_DRIVER")
	if os.Getenv("DB_DRIVER") == "" {
		env.Driver = utils.GetEnv("DB_DRIVER", "mysql")
	}

	return &env
}
