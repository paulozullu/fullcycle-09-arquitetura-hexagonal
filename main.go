package main

import (
	"database/sql"

	"github.com/paulozullu/09-arquitetura-hexagonal/adapters/db"
	"github.com/paulozullu/09-arquitetura-hexagonal/application"
)

func main() {
	Db, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := db.NewProductDb(Db)

	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}
