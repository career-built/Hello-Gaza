package api

import (
	"example/baseProject/product"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductRouter struct {
	productManager product.ProductManager
}

func NewProductRouter(productManager product.ProductManager) *ProductRouter {
	return &ProductRouter{
		productManager: productManager,
	}
}

func (obj *ProductRouter) CreateProduct(c echo.Context) error {
	//Request handeling
	product := new(product.Product)

	if err := c.Bind(product); err != nil {
		fmt.Printf("Error While Binding the product\n")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//Service invoc
	if err := obj.productManager.Add(product, ""); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Respond handeling
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Product created successfully",
		"name":    product.NAME,
	})
}
func (obj *ProductRouter) GetProductByID(c echo.Context) error {

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
	product, err := obj.productManager.GetByID(productID)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Unaple to Fetch product from DB",
		})
	}
	if product == nil {
		return c.JSON(http.StatusOK, map[string]string{
			"INFO": "Product Not Found",
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
