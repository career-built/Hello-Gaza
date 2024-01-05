package main

import (
	"fmt"
	"hello-gaza/controller"
	"hello-gaza/database"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hiiiiiiiiiiiiiiiiiiii2")

	database.InitDB()
	defer database.CloseDB()
	fmt.Println("Hiiiiiiiiiiiiiiiiiiii22222")
	// Insert a product into the database
	e := echo.New()
	e.POST("/product/create", controller.CreateProduct)
	e.GET("/product/:id", controller.GetProductByID)
	e.Logger.Fatal(e.Start(":3030"))

	fmt.Println("Selected Product:")
}
