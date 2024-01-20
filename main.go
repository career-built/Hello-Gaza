package main

import (
	"example/baseProject/api"
	"example/baseProject/database"
	"example/baseProject/messageBroker"
	"example/baseProject/product"

	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting Base")
	// define the base dbConnector
	dbConnector := database.NewPostgres()
	if dbConnector == nil {
		log.Fatal("can't connect to database")
	}
	defer dbConnector.CloseDB()
	// define the base Message Broker
	broker, err := messageBroker.NewRabbitMQBroker("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer broker.Close()

	//inject the Database to the base Db interface
	//inject the messege to the base broker interface
	productMgr := product.NewProductService(dbConnector, broker)

	//inject the product feature to the base manager interface
	productRouter := api.NewProductRouter(productMgr)

	e := echo.New()
	e.POST("/product/create", productRouter.CreateProduct)
	e.GET("/product/:id", productRouter.GetProductByID)

	e.Logger.Fatal(e.Start(":3030"))

}
