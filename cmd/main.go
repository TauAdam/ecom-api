package main

import (
	"github.com/TauAdam/ecom-api/cmd/api"
	"github.com/TauAdam/ecom-api/config"
	my_sql "github.com/TauAdam/ecom-api/internal/storage/mysql"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db := my_sql.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		DBName:               config.Envs.DBName,
		Addr:                 config.Envs.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
