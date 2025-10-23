package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

func main() {
	router := gin.Default()
	//initDB()

	router.POST("/login", login)

	allPCRUDfuncs(router, "/customers", &customers, reflect.TypeOf(Customer{}))
	allPCRUDfuncs(router, "/markets", &markets, reflect.TypeOf(Market{}))
	allPCRUDfuncs(router, "/employees", &employees, reflect.TypeOf(Employee{}))
	allPCRUDfuncs(router, "/posts", &posts, reflect.TypeOf(Post{}))
	allPCRUDfuncs(router, "/orders", &orders, reflect.TypeOf(Order{}))
	allPCRUDfuncs(router, "/providers", &providers, reflect.TypeOf(Provider{}))
	allCRUDfuncs(router, "/dishes", &dishes, reflect.TypeOf(Dish{}))
	allCRUDfuncs(router, "/ingredients", &ingredients, reflect.TypeOf(Ingredient{}))
	allCRUDfuncs(router, "/categories", &categories, reflect.TypeOf(Category{}))

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server running at port 8080")
}
