package my_sql

import (
	"database/sql"
	"log"
)
import "github.com/go-sql-driver/mysql"

func NewMySQLStorage(cfg mysql.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping mysql: %v", err)
	}
	log.Println("mysql storage is ready")
	return db
}
