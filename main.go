package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/customers/:id", getCustomerID)
	router.GET("/customers", getCustomers)
	router.GET("/markets", getMarkets)
	router.GET("/markets/:id", getMarketID)
	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", getEmployeeID)
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostID)
	router.GET("/orders", getOrders)
	router.GET("/orders/:id", getOrderID)
	router.GET("/providers", getProviders)
	router.GET("/providers/:id", getProviderID)
	router.GET("/dishes", getDishes)
	router.GET("/dishes/:id", getDishID)
	router.GET("/ingredients", getIngredients)
	router.GET("/ingredients/:id", getIngredientID)
	router.GET("/categories", getCategories)
	router.GET("/categories/:id", getCategoryID)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server running at port 8080")
}
