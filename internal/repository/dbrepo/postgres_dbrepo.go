package dbrepo

import (
	"context"
	"database/sql"
	"project-control-backend/internal/models"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

// Connection returns the database connection
func (p *PostgresDBRepo) Connection() *sql.DB {
	return p.DB
}

const dbTimeout = time.Second * 3

//GetVendors returns all vendors from the database
func (p *PostgresDBRepo) GetVendors() ([]*models.Vendor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	query := `SELECT id, name, email, phone, street, city, state, zip,  created_at, updated_at FROM vendor ORDER BY name`
	
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var vendors []*models.Vendor
	
	for rows.Next() {
		var vendor models.Vendor
		err := rows.Scan(&vendor.ID, &vendor.Name, &vendor.Email, &vendor.Phone, &vendor.Street, &vendor.City, &vendor.State, &vendor.Zip, &vendor.CreatedAt, &vendor.UpdatedAt)
		if err != nil {
			return nil, err
		}
		vendors = append(vendors, &vendor)
	}
	
	return vendors, nil
}

//GetVendorByID returns a single vendor from the database
func (p *PostgresDBRepo) GetVendorByID(id int) (*models.Vendor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	query := `SELECT id, name, email, phone, street, city, state, zip,  created_at, updated_at FROM vendor WHERE id = $1`
	
	row := p.DB.QueryRowContext(ctx, query, id)
	
	var vendor models.Vendor
	
	err := row.Scan(&vendor.ID, &vendor.Name, &vendor.Email, &vendor.Phone, &vendor.Street, &vendor.City, &vendor.State, &vendor.Zip, &vendor.CreatedAt, &vendor.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return &vendor, nil
}

//GetProducts returns all products from the database
func (p *PostgresDBRepo) GetProducts() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	query := `SELECT  id, name, price, created_at, updated_at FROM products ORDER BY name`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []*models.Product
	
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	
	return products, nil
}

//GetProductByID returns a single product from the database
func (p *PostgresDBRepo) GetProductByID(id int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	query := `SELECT id, name, price, created_at, updated_at FROM products WHERE id = $1`
	
	row := p.DB.QueryRowContext(ctx, query, id)
	
	var product models.Product
	
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return &product, nil
}
