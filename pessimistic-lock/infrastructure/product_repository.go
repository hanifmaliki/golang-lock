package infrastructure

import (
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) UpdateProductPrice(id int, newPrice float64) error {
	// Begin a new transaction
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	// Ensure the transaction is rolled back if there's an error
	defer tx.Rollback()

	// Lock the product for update within the transaction
	var price float64
	query := `SELECT price FROM products WHERE id = $1 FOR UPDATE`
	err = tx.QueryRow(query, id).Scan(&price)
	if err != nil {
		return fmt.Errorf("error locking product: %v", err)
	}

	// Update the price within the same transaction
	updateQuery := `UPDATE products SET price = $1 WHERE id = $2`
	_, err = tx.Exec(updateQuery, newPrice, id)
	if err != nil {
		return fmt.Errorf("error updating product price: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	fmt.Printf("Updated product price to %.2f\n", newPrice)
	return nil
}
