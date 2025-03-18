package db_test

import (
	"database/sql"
	"go-hexagonal/adapters/db"
	"go-hexagonal/application"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products 
	(
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("abc", "product test", 0, "disabled");`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	ProductDb := db.NewProductDb(Db)
	Product, err := ProductDb.Get("abc")

	require.Nil(t, err)

	require.Equal(t, "abc", Product.GetID())
	require.Equal(t, "product test", Product.GetName())
	require.Equal(t, 0.0, Product.GetPrice())
	require.Equal(t, "disabled", Product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	ProductDb := db.NewProductDb(Db)
	Product := application.NewProduct()
	Product.Name = "product test"
	Product.Price = 0

	ProductResult, err := ProductDb.Save(Product)

	require.Nil(t, err)
	require.Equal(t, Product.Name, ProductResult.GetName())
	require.Equal(t, Product.Price, ProductResult.GetPrice())
	require.Equal(t, Product.Status, ProductResult.GetStatus())

	Product.Price = 10
	Product.Enable()
	ProductResult, err = ProductDb.Save(Product)

	require.Nil(t, err)
	require.Equal(t, Product.Price, ProductResult.GetPrice())
	require.Equal(t, Product.Status, ProductResult.GetStatus())
}
