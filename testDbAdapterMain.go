package main

import (
	"database/sql"
	dbAdapter "go-hexagonal/adapters/db"
	"go-hexagonal/application"

	_ "github.com/mattn/go-sqlite3"
)

func testDbAdapterMain() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := dbAdapter.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product 1", 10.0)

	productService.Enable(product)
}
