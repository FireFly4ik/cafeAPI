package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	allCRUDfuncs(router, "/customers", &customers)
	allCRUDfuncs(router, "/markets", &markets)
	allCRUDfuncs(router, "/employees", &employees)
	allCRUDfuncs(router, "/posts", &posts)
	allCRUDfuncs(router, "/orders", &orders)
	allCRUDfuncs(router, "/providers", &providers)
	allCRUDfuncs(router, "/dishes", &dishes)
	allCRUDfuncs(router, "/ingredients", &ingredients)
	allCRUDfuncs(router, "/categories", &categories)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server running at port 8080")
}
