package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// все функции для всех путей
func allCRUDfuncs[T any](r *gin.Engine, path string, data *[]T) {
	r.GET(path, func(c *gin.Context) {
		c.JSON(http.StatusOK, *data)
	})

	r.GET(path+"/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, item := range *data {
			if id == getItemID(item) {
				c.JSON(http.StatusOK, item)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
	})

	r.DELETE(path+"/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for ind, item := range *data {
			if id == getItemID(item) {
				*data = append((*data)[:ind], (*data)[ind+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	})

	r.POST(path, func(c *gin.Context) {
		var item T
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
			return
		}
		*data = append(*data, item)
		c.JSON(http.StatusCreated, gin.H{"message": "Item added"})
	})

	r.PUT(path+"/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var newItem T
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
			return
		}
		for ind, item := range *data {
			if id == getItemID(item) {
				(*data)[ind] = newItem
				c.JSON(http.StatusOK, gin.H{"message": "Item updated"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
	})
}

func getItemID[T any](item T) int {
	switch v := any(item).(type) {
	case Customer:
		return v.CustomerID
	case Employee:
		return v.EmployeeID
	case Market:
		return v.MarketID
	case Post:
		return v.PostID
	case Order:
		return v.OrderID
	case Provider:
		return v.ProviderID
	case Ingredient:
		return v.IngredientID
	case Dish:
		return v.DishID
	case Category:
		return v.CategoryID
	}
	return 0
}
