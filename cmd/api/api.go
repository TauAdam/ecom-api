package api

import (
	"database/sql"
	"github.com/TauAdam/ecom-api/internal/modules/cart"
	"github.com/TauAdam/ecom-api/internal/modules/products"
	"github.com/TauAdam/ecom-api/internal/modules/user"
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

	userStore := user.NewUserStore(s.db)
	userController := user.NewHandler(userStore)
	userController.InitRoutes(subrouter)

	productStore := products.NewProductsStore(s.db)
	productHandler := products.NewHandler(productStore)
	productHandler.InitRoutes(subrouter)

	cartStore := cart.NewCartStore(s.db)
	cartHandler := cart.NewHandler(cartStore, productStore, userStore)
	cartHandler.InitRoutes(subrouter)

	log.Printf("Server is listening on %s", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}
