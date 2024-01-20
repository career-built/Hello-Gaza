package product

import (
	"example/baseProject/database"
	"fmt"
	"log"
	"strconv"
)

type ProductService struct {
	dbConnector database.DbConnector
}

func NewProductService(connector database.DbConnector) *ProductService {
	return &ProductService{
		dbConnector: connector,
	}
}

func (obj *ProductService) Add(product *Product) error {
	priceStr := strconv.Itoa(product.PRICE)

	keys := []string{"name", "price"}
	vals := []string{
		product.NAME,
		priceStr,
	}
	err := obj.dbConnector.Insert("product", keys, vals)
	if err != nil {
		log.Println("Product not inserted", err)
		return err
	}
	return nil
}

func (obj *ProductService) GetByID(id int) *Product {
	row, err := obj.dbConnector.SelectById("product", id)
	if err != nil {
		log.Printf("Error while retrive the product with %d , Erorr:,%s ", id, err)
		return nil
	}

	fmt.Println(row)
	return nil
}
