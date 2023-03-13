package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {
	connection := "user=postgres dbname=go_store password=XXXXXX host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()

	productsSQL, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for productsSQL.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsSQL.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "index", products)
	defer db.Close()
}
