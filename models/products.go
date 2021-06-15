package models

import "store/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Stock       int
}

func GetAllProducts() []Product {
	conn := db.DatabaseConector()

	selectAllProducts, err := conn.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, stock int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &stock)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Stock = stock

		products = append(products, p)
	}

	defer conn.Close()
	return products
}

func NewProduct(name, description string, price float64, stock int) {
	conn := db.DatabaseConector()

	insertDB, err := conn.Prepare("insert into products(name, description, price, stock) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertDB.Exec(name, description, price, stock)

	defer conn.Close()
}

func DeleteProduct(id string) {
	conn := db.DatabaseConector()

	deleteDB, err := conn.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteDB.Exec(id)

	defer conn.Close()
}

func GetProduct(id string) Product {
	conn := db.DatabaseConector()

	productDB, err := conn.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	for productDB.Next() {
		var id, stock int
		var name, description string
		var price float64

		err = productDB.Scan(&id, &name, &description, &price, &stock)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Stock = stock
	}
	defer conn.Close()
	return p
}

func UpdateProduct(id int, name, description string, price float64, stock int) {
	conn := db.DatabaseConector()

	updateDB, err := conn.Prepare("update products set name=$1, description=$2, price=$3, stock=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateDB.Exec(name, description, price, stock, id)

	defer conn.Close()
}
