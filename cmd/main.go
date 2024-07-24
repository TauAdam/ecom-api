package main

import (
	"github.com/TauAdam/ecom-api/cmd/api"
	"github.com/TauAdam/ecom-api/config"
	my_sql "github.com/TauAdam/ecom-api/internal/storage/mysql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	cfg := config.NewConfig()

	db, err := my_sql.NewMySQLStorage(mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		DBName:               cfg.DBName,
		Addr:                 cfg.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
