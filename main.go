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
	router.DELETE("/customers/:id", deleteCustomerID)
	router.DELETE("/markets/:id", deleteMarketID)
	router.DELETE("/employees/:id", deleteEmployeeID)
	router.DELETE("/posts/:id", deletePostID)
	router.DELETE("/orders/:id", deleteOrderID)
	router.DELETE("/providers/:id", deleteProviderID)
	router.DELETE("/dishes/:id", deleteDishID)
	router.DELETE("/ingredients/:id", deleteIngredientID)
	router.DELETE("/categories/:id", deleteCategoryID)
	router.POST("/customers", addCustomer)
	router.POST("/markets", addMarket)
	router.POST("/posts", addPost)
	router.POST("/categories", addCategory)
	router.POST("/ingredients", addIngredient)
	router.POST("/dishes", addDish)
	router.POST("/providers", addProvider)
	router.POST("/orders", addOrder)
	router.POST("/employees", addEmployee)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server running at port 8080")
}
