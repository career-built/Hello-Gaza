package database

import (
	"database/sql"
	"fmt"
	"log"

	"hello-gaza/models"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres_container"
	port     = 5433
	user     = "postgres"
	password = "1234"
	dbname   = "test"
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func CloseDB() {
	db.Close()
	fmt.Println("Closed the database connection")
}

func InsertProduct(product *models.Product) error {
	query := "INSERT INTO product (name,price) VALUES ($1,$2)"
	fmt.Printf("Insert {$1}{$2}"+product.NAME, product.PRICE)
	_, err := db.Query(query, product.NAME, product.PRICE)
	if err != nil {
		log.Println("Error inserting product:", err)
		return err
	}

	fmt.Println("Product inserted successfully")
	return nil
}

func GetProductByID(id int) (models.Product, error) {
	var product models.Product
	query := "SELECT id, name,price FROM product WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&product.ID, &product.NAME, &product.PRICE)
	if err != nil {
		log.Println("Error getting product:", err)
		return models.Product{}, err
	}

	return product, nil
}
