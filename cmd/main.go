package main

import (
	"github.com/TauAdam/ecom-api/cmd/api"
	"github.com/TauAdam/ecom-api/config"
	my_sql "github.com/TauAdam/ecom-api/internal/storage/mysql"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.NewConfig()

	db := my_sql.NewMySQLStorage(mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		DBName:               cfg.DBName,
		Addr:                 cfg.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
