package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getCustomerID(c *gin.Context) {
	id := c.Param("id")
	for _, customer := range customers {
		if id == fmt.Sprintf("%d", customer.CustomerID) {
			c.JSON(http.StatusOK, customer)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Cafe not found"})
	return
}

func getCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, customers)
}

func getMarketID(c *gin.Context) {
	id := c.Param("id")
	for _, market := range markets {
		if id == fmt.Sprintf("%d", market.MarketID) {
			c.JSON(http.StatusOK, market)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Market not found"})
}

func getMarkets(c *gin.Context) {
	c.JSON(http.StatusOK, markets)
}

func getEmployeeID(c *gin.Context) {
	id := c.Param("id")
	for _, employee := range employees {
		if id == fmt.Sprintf("%d", employee.EmployeeID) {
			c.JSON(http.StatusOK, employee)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func getEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

func getPostID(c *gin.Context) {
	id := c.Param("id")
	for _, post := range posts {
		if id == fmt.Sprintf("%d", post.PostID) {
			c.JSON(http.StatusOK, post)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func getOrderID(c *gin.Context) {
	id := c.Param("id")
	for _, order := range orders {
		if id == fmt.Sprintf("%d", order.OrderID) {
			c.JSON(http.StatusOK, order)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func getOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

func getProviderID(c *gin.Context) {
	id := c.Param("id")
	for _, provider := range providers {
		if id == fmt.Sprintf("%d", provider.ProviderID) {
			c.JSON(http.StatusOK, provider)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Provider not found"})
}

func getProviders(c *gin.Context) {
	c.JSON(http.StatusOK, providers)
}

func getDishID(c *gin.Context) {
	id := c.Param("id")
	for _, dish := range dishes {
		if id == fmt.Sprintf("%d", dish.DishID) {
			c.JSON(http.StatusOK, dish)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Dish not found"})
}

func getDishes(c *gin.Context) {
	c.JSON(http.StatusOK, dishes)
}

func getIngredientID(c *gin.Context) {
	id := c.Param("id")
	for _, ingredient := range ingredients {
		if id == fmt.Sprintf("%d", ingredient.IngredientID) {
			c.JSON(http.StatusOK, ingredient)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Ingredient not found"})
}

func getIngredients(c *gin.Context) {
	c.JSON(http.StatusOK, ingredients)
}

func getCategoryID(c *gin.Context) {
	id := c.Param("id")
	for _, category := range categories {
		if id == fmt.Sprintf("%d", category.CategoryID) {
			c.JSON(http.StatusOK, category)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

func getCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}
