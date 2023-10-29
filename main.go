package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(db)
	}

	productUseCase := NewProductUseCase(db)

	product, err := productUseCase.GetProduct(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.Name)
}
