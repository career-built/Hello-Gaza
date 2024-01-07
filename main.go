package main

import (
	"database/sql"
	"fmt"
	"hello-gaza/api"
	"hello-gaza/controller"
	"hello-gaza/db"
	"hello-gaza/product"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "test"
)

func main() {

	fmt.Println("Hiiiiiiiiiiiiiiiiiiii2")
	//
	//database.InitDB()
	//defer database.CloseDB()
	fmt.Println("Hiiiiiiiiiiiiiiiiiiii22222")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	//err = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}

	sqlBuilder := db.NewSqlBuilder()
	productMgr := product.NewManager(dbConn, sqlBuilder)
	productRouter := api.NewProductRouter(productMgr)

	// Insert a product into the database
	e := echo.New()

	//e.POST("/product/create", controller.CreateProduct)
	e.POST("/product/create", productRouter.CreateProduct)
	e.GET("/product/:id", controller.GetProductByID)
	e.Logger.Fatal(e.Start(":3030"))

	fmt.Println("Selected Product:")
}
