package main

import (
	"fmt"

	"jeffpalmeri.com/stored-proc/m/src/config"
	"jeffpalmeri.com/stored-proc/m/src/models"
)

func main() {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println("An error occured: ", err)
	}

	productModel := models.ProductModel{
		Db: db,
	}
	fmt.Println("Product List")

	products, err2 := productModel.Search(250, 250)

	if err2 != nil {
		fmt.Println("An error occured 2: ", err2)
	}

	fmt.Print("Products: ", len(products), "\n")

	for _, product := range products {
		fmt.Println("Id:", product.Id)
		fmt.Println("Name:", product.Name)
		fmt.Println("Price:", product.Price)
		fmt.Println("Quantity:", product.Quantity)
		fmt.Println("Status:", product.Status)
		fmt.Println("----------------------------")
	}
}