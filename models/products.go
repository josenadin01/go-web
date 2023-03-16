package models

import "go-web/database"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func GetAllProducts() []Product {
	db := database.DatabaseConnection()

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
	defer db.Close()
	return products
}
