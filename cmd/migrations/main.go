package main

import (
	"errors"
	my_sql "github.com/TauAdam/ecom-api/internal/storage/mysql"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func main() {
	db := my_sql.NewMySQLStorage(mysqlCfg.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		DBName:               cfg.DBName,
		Addr:                 cfg.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrations", "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
}
