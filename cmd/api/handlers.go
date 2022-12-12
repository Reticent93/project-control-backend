package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"project-control-backend/internal/helpers"
	"strconv"
)

//Home handler for the home page
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Welcome to the Go Project API",
		Version: "1.0.0",
	}
	_ = app.jsonResponse(w, http.StatusOK, payload)
}

//GetVendors handler for the home page
func (app *application) GetVendors(w http.ResponseWriter, r *http.Request) {
	vendors, err := app.DB.GetVendors()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_ = app.jsonResponse(w, http.StatusOK, vendors)
}

//GetVendor handler for the home page
func (app *application) GetVendor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	vendorID, err := strconv.Atoi(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	
	vendor, err := app.DB.GetVendorByID(vendorID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_ = app.jsonResponse(w, http.StatusOK, vendor)
}

//GetProducts handler for the home page
func (app *application) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := app.DB.GetProducts()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_ = app.jsonResponse(w, http.StatusOK, products)
}

//GetProduct handler for the home page
func (app *application) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	
	product, err := app.DB.GetProductByID(productID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_ = app.jsonResponse(w, http.StatusOK, product)
}
