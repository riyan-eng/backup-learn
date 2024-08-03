package main

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/blockloop/tea"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

type PostBreweryRequest struct {
	Name string `json:"name" validate:"required"`
	City string `json:"city" validate:"required"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(httprate.Limit(
		2,
		1*time.Second,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "some specific response here", http.StatusTooManyRequests)
		}),
	))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		lala := PostBreweryRequest{
			Name: "riyan",
			City: "Batang",
		}
		// panic("mengapa")
		jsonLala, err := json.Marshal(lala)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonLala)
	})

	r.Get("/lala", tea.Handler(GetBrewery))
	http.ListenAndServe(":3000", r)
}

func GetBrewery(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	lala := PostBreweryRequest{
		Name: "riyan",
		City: "Batang",
	}

	return 200, lala
}
