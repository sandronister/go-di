package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/go-di/product"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(db)
	}

	productRepository := product.NewProductRepository(db)
	productUseCase := product.NewProductUseCase(productRepository)

	product, err := productUseCase.GetProduct(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.Name)
}
