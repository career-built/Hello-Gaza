package main

import (
	"example/baseProject/api"
	"example/baseProject/database"
	"example/baseProject/product"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting Base")
	dbConnector := database.NewPostgres()
	if dbConnector == nil {
		log.Fatal("can't connect to database")
	}
	defer dbConnector.CloseDB()

	//Path the Database to the base Db interface
	productMgr := product.NewProductService(dbConnector)

	//Path the product feature to the base manger interface
	productRouter := api.NewProductRouter(productMgr)

	e := echo.New()
	e.POST("/product/create", productRouter.CreateProduct)
	e.GET("/product/:id", productRouter.GetProductByID)

	e.Logger.Fatal(e.Start(":3030"))

}
