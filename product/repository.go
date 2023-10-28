package product

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r ProductRepository) GetProduct(id int) (Product, error) {
	return Product{ID: id, Name: "Product Name"}, nil
}
