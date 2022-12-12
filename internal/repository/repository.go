package repository

import (
	"database/sql"
	"project-control-backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetVendors() ([]*models.Vendor, error)
	GetVendorByID(id int) (*models.Vendor, error)
	GetProducts() ([]*models.Product, error)
	GetProductByID(id int) (*models.Product, error)
}
