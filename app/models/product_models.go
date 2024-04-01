package models

import (
	"belajar-golang/connection"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetAllProducts() ([]Product, error) {
	rows, err := connection.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		var createdAt, updatedAt string
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		// Konversi string ke time.Time
		p.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		p.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
		products = append(products, p)
	}
	return products, nil
}


func CreateProduct(p Product) (*Product, error) {
	result, err := connection.DB.Exec("INSERT INTO products (name, description, price, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		p.Name, p.Description, p.Price, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	// Mendapatkan ID produk yang baru saja ditambahkan
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	// Mengembalikan produk dengan ID yang baru saja ditambahkan
	return &Product{
		ID:          uint(id),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func DeleteProduct(id uint) error {
	_, err := connection.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func UpdateProduct(p Product) error {
    result, err := connection.DB.Exec("UPDATE products SET name = ?, description = ?, price = ?, updated_at = ? WHERE id = ?", p.Name, p.Description, p.Price, time.Now(), p.ID)
    if err != nil {
        return err
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return fmt.Errorf("product with ID %d not found", p.ID)
    }
    return nil
}
func GetProductById(id uint) (*Product, error) {
	row := connection.DB.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var p Product
	var createdAt, updatedAt string
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	// Konversi string ke time.Time
	p.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	p.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
	return &p, nil
}
