package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	
	mux.Get("/", app.Home)
	mux.Get("/vendors", app.GetVendors)
	mux.Get("/vendors/{id}", app.GetVendor)
	mux.Get("/products", app.GetProducts)
	mux.Get("/products/{id}", app.GetProduct)
	
	return mux
}
