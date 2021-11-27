package models

import (
	"database/sql"

	"jeffpalmeri.com/stored-proc/m/src/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) Search(min, max float64) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("call findBetween(?, ?)", min, max)

	if err != nil {
		return nil, err
	}

	products := []entities.Product{}

	for rows.Next() {
		var id int64
		var name string
		var price float32
		var quantity int
		var status bool
		err2 := rows.Scan(&id, &name, &price, &quantity, &status)

		if err2 != nil {
			return nil, err2
		}

		// product := entities.Product{id, name, price, quantity, status}
		product := entities.Product{
			Id: id, 
			Name: name, 
			Price: price, 
			Quantity: quantity, 
			Status: status,
		}
		products = append(products, product)
	}
	return products, nil
}
