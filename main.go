package main

import (
	"log"
	"net/http"

	"github.com/Scrackc/api-crud-go/auth"
	"github.com/Scrackc/api-crud-go/handler"
	"github.com/Scrackc/api-crud-go/storage"
)

func main() {

	err := auth.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudieron cargar los certificados: %v", err)
	}

	store := storage.NewMemory()

	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteAuth(mux, &store)
	log.Println("Servidor iniciado en el puerto :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("error en el servidor: %v\n", err)
	}
}
