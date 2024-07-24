package api

import (
	"database/sql"
	"log"
	"net/http"
)
import "github.com/gorilla/mux"

type Server struct {
	addr string
	db   *sql.DB
}

func (s Server) Run() interface{} {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Printf("Server is listening on %s", s.addr)
	return http.ListenAndServe(s.addr, subrouter)
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}
