package handler

import (
	"net/http"

	"github.com/Scrackc/api-crud-go/middleware"
)

// RoutePerson .
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", middleware.Authentication(h.create))
	mux.HandleFunc("/v1/persons/get-all", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/get-by-id", h.getByID)

}

func RouteAuth(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)

	mux.HandleFunc("/v1/login", h.login)
}
