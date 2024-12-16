package infrastructure

import (
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) UpdateProductPrice(id int, newPrice float64) error {
	query := `UPDATE products SET price = $1 WHERE id = $2`
	_, err := r.db.Exec(query, newPrice, id)
	if err != nil {
		return fmt.Errorf("error updating product price: %v", err)
	}

	fmt.Printf("Updated product price to %.2f\n", newPrice)
	return nil
}
