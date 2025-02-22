package main

import (
	"encoding/json"
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
	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
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

func deleteCustomerID(c *gin.Context) {
	id := c.Param("id")
	for ind, customer := range customers {
		if id == fmt.Sprintf("%d", customer.CustomerID) {
			customers = append(customers[:ind], customers[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
	return
}

func deleteMarketID(c *gin.Context) {
	id := c.Param("id")
	for ind, market := range markets {
		if id == fmt.Sprintf("%d", market.MarketID) {
			markets = append(markets[:ind], markets[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Market deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Market not found"})
	return
}

func deleteEmployeeID(c *gin.Context) {
	id := c.Param("id")
	for ind, employee := range employees {
		if id == fmt.Sprintf("%d", employee.EmployeeID) {
			employees = append(employees[:ind], employees[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
	return
}

func deletePostID(c *gin.Context) {
	id := c.Param("id")
	for ind, post := range posts {
		if id == fmt.Sprintf("%d", post.PostID) {
			posts = append(posts[:ind], posts[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
	return
}

func deleteOrderID(c *gin.Context) {
	id := c.Param("id")
	for ind, order := range orders {
		if id == fmt.Sprintf("%d", order.OrderID) {
			orders = append(orders[:ind], orders[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
	return
}

func deleteProviderID(c *gin.Context) {
	id := c.Param("id")
	for ind, provider := range providers {
		if id == fmt.Sprintf("%d", provider.ProviderID) {
			providers = append(providers[:ind], providers[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Provider deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Provider not found"})
	return
}

func deleteDishID(c *gin.Context) {
	id := c.Param("id")
	for ind, dish := range dishes {
		if id == fmt.Sprintf("%d", dish.DishID) {
			dishes = append(dishes[:ind], dishes[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Dish deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Dish not found"})
	return
}

func deleteIngredientID(c *gin.Context) {
	id := c.Param("id")
	for ind, ingredient := range ingredients {
		if id == fmt.Sprintf("%d", ingredient.IngredientID) {
			ingredients = append(ingredients[:ind], ingredients[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Ingredient deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Ingredient not found"})
	return
}

func deleteCategoryID(c *gin.Context) {
	id := c.Param("id")
	for ind, category := range categories {
		if id == fmt.Sprintf("%d", category.CategoryID) {
			categories = append(categories[:ind], categories[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
	return
}

func addCustomer(c *gin.Context) {
	var customer Customer
	err := json.NewDecoder(c.Request.Body).Decode(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	customer.CustomerID = len(customers) + 1
	customers = append(customers, customer)
	c.JSON(http.StatusOK, gin.H{"message": "Customer added"})
}

func addMarket(c *gin.Context) {
	var market Market
	err := json.NewDecoder(c.Request.Body).Decode(&market)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	market.MarketID = len(markets) + 1
	markets = append(markets, market)
	c.JSON(http.StatusOK, gin.H{"message": "Market added"})
}

func addPost(c *gin.Context) {
	var post Post
	err := json.NewDecoder(c.Request.Body).Decode(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	post.PostID = len(posts) + 1
	posts = append(posts, post)
	c.JSON(http.StatusOK, gin.H{"message": "Post added"})
}

func addCategory(c *gin.Context) {
	var category Category
	err := json.NewDecoder(c.Request.Body).Decode(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	category.CategoryID = len(categories) + 1
	categories = append(categories, category)
	c.JSON(http.StatusOK, gin.H{"message": "Category added"})
}

func addIngredient(c *gin.Context) {
	var ingredient Ingredient
	err := json.NewDecoder(c.Request.Body).Decode(&ingredient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	ingredient.IngredientID = len(ingredients) + 1
	ingredients = append(ingredients, ingredient)
	c.JSON(http.StatusOK, gin.H{"message": "Ingredient added"})
}

func addDish(c *gin.Context) {
	var dish Dish
	err := json.NewDecoder(c.Request.Body).Decode(&dish)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	dish.DishID = len(dishes) + 1
	dishes = append(dishes, dish)
	c.JSON(http.StatusOK, gin.H{"message": "Dish added"})
}

func addProvider(c *gin.Context) {
	var provider Provider
	err := json.NewDecoder(c.Request.Body).Decode(&provider)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	provider.ProviderID = len(providers) + 1
	providers = append(providers, provider)
	c.JSON(http.StatusOK, gin.H{"message": "Provider added"})
}

func addOrder(c *gin.Context) {
	var order Order
	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	order.OrderID = len(orders) + 1
	orders = append(orders, order)
	c.JSON(http.StatusOK, gin.H{"message": "Order added"})
}

func addEmployee(c *gin.Context) {
	var employee Employee
	err := json.NewDecoder(c.Request.Body).Decode(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
		return
	}
	employee.EmployeeID = len(employees) + 1
	employees = append(employees, employee)
	c.JSON(http.StatusOK, gin.H{"message": "Employee added"})
}
