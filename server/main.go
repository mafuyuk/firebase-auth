package main

import (
	"net/http"
	"time"
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"fmt"
)

// curl http://localhost:8080/v1/hello
func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/public", public)
	r.Get("/private", private)

	http.ListenAndServe(":8000", r)
}


func public(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	rqqID := middleware.GetReqID(ctx)
	log.Println(rqqID)

	w.WriteHeader(200)
	w.Write([]byte("hello public!\n"))

}

func private(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	rqqID := middleware.GetReqID(ctx)
	log.Println(rqqID)

	w.WriteHeader(200)
	w.Write([]byte("hello private!\n"))

}
