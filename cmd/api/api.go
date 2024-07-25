package api

import (
	"database/sql"
	"github.com/TauAdam/ecom-api/internal/controllers/user"
	"github.com/TauAdam/ecom-api/internal/store"
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

	userStore := store.NewUserStore(s.db)
	userController := user.NewHandler(userStore)
	userController.InitRoutes(subrouter)

	log.Printf("Server is listening on %s", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}
