package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Scrackc/api-crud-go/auth"
)

type FuncHandler func(w http.ResponseWriter, r *http.Request)

// Log .
func Log(f FuncHandler) FuncHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		log.Printf("petición: %q, método: %q, duration:%v", r.URL.Path, r.Method, time.Since(start).Milliseconds())
	}
}

// Authentication .
func Authentication(f FuncHandler) FuncHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := auth.ValidateToken(token)
		fmt.Println(err)
		if err != nil {
			forbiden(w, r)
			return
		}
		f(w, r)
	}
}

func forbiden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"message": "No authenticado"}`))
}
