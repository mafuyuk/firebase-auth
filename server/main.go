package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// curl http://localhost:8080/v1/hello
func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/hello", handler)
	})

	http.ListenAndServe(":8080", r)
}


func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rqqID := middleware.GetReqID(ctx)
	logger := middleware.GetLogEntry(r)
	logger.Write(200, 1, time.Second)

	w.Write([]byte(`{"request_id":"`+rqqID+`"}`))
}
