package api

import (
	"github.com/gorilla/mux"
	"github.com/avijit-dhaliwal/RESTfulAPI/internal/api"
	"github.com/avijit-dhaliwal/RESTfulAPI/internal/db"
)

func NewRouter(database *db.Database) *mux.Router {
	router := mux.NewRouter()
	handler := &Handler{DB: database}

	router.HandleFunc("/items", auth.Middleware(handler.GetItems)).Methods("GET")
	router.HandleFunc("/items", auth.Middleware(handler.CreateItem)).Methods("POST")
	router.HandleFunc("/items/{id}", auth.Middleware(handler.GetItem)).Methods("GET")
	router.HandleFunc("/items/{id}", auth.Middleware(handler.UpdateItem)).Methods("PUT")
	router.HandleFunc("/items/{id}", auth.Middleware(handler.DeleteItem)).Methods("DELETE")

	return router
}