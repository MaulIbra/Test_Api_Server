package config

import (
	"database/sql"
	"fmt"
	logs "github.com/MaulIbra/logs_module"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(env *Env) *sql.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.SchemaName)
	db, _ := sql.Open(env.Driver, dataSource)

	if err := db.Ping(); err != nil {
		panic(err)
		logs.ErrorLogger.Println(err)
		return nil
	}
	return db
}
