package infrastructure

import (
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) UpdateProductPrice(id int, newPrice float64) error {
	// Get the current version of the product
	var version int
	err := r.db.QueryRow("SELECT version FROM products WHERE id = $1", id).Scan(&version)
	if err != nil {
		return fmt.Errorf("error fetching product version: %v", err)
	}

	// Attempt to update the product price directly, with version check
	query := `UPDATE products SET price = $1, version = version + 1 WHERE id = $2 AND version = $3`
	result, err := r.db.Exec(query, newPrice, id, version)
	if err != nil {
		return fmt.Errorf("error updating product price: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("update failed, version mismatch")
	}

	fmt.Printf("Updated product price to %.2f\n", newPrice)
	return nil
}
