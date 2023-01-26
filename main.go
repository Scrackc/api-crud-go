package main

import (
	"log"
	"net/http"

	"github.com/Scrackc/api-crud-go/handler"
	"github.com/Scrackc/api-crud-go/storage"
)

func main() {
	store := storage.NewMemory()

	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	log.Println("Servidor iniciado en el puerto :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("error en el servidor: %v\n", err)
	}
}
