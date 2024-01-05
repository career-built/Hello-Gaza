package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"hello-gaza/database"
	"hello-gaza/models"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Product struct {
	Price float64 `json:"price"`
}

// ---------- Create Product -------------
func CreateProduct(c echo.Context) error {
	// Bind the request body to the Product struct
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		fmt.Printf("Error While Binding the product\n")
		return err
	}
	// Save the product to the database
	err := database.InsertProduct(product)
	if err != nil {
		fmt.Printf("Error while saving the product to the database\n")
		return err
	}
	// Respond with a JSON message
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Product created successfully",
		"name":    product.NAME,
	})
}

// ---------- Get Product -------------
func GetProductByID(c echo.Context) error {

	productIDStr := c.Param("id")

	// Check if the productID is a valid integer
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	// Fetch the corresponding product from the database
	fmt.Printf("new Requested product ID: %d\n", productID)
	product := new(models.Product)
	*product, err = database.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Unaple to Fetch product from DB",
		})
	}

	// Respond with a JSON message
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Product fetched successfully",
		"id":      strconv.Itoa(product.ID),
		"name":    product.NAME,
		"price":   strconv.Itoa(product.PRICE),
	})
}
