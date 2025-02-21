package main

import "time"

type Customer struct {
	CustomerID int    `json:"customerID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"title"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Bonuses    int    `json:"bonuses"`
}

type Market struct {
	MarketID  int    `json:"marketID"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	StartHour int    `json:"startHour"`
	EndHour   int    `json:"endHour"`
	Phone     string `json:"phone"`
}

type Employee struct {
	EmployeeID int    `json:"employeeID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Passport   string `json:"passport"`
	Phone      string `json:"phone"`
	PostID     int    `json:"postID"`
}

type Post struct {
	PostID int    `json:"postID"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type Order struct {
	OrderID    int       `json:"orderID"`
	DateTime   time.Time `json:"dateTime"`
	EmployeeID int       `json:"employeeID"`
	Closed     bool      `json:"closed"`
	CustomerID int       `json:"customerID"`
}

type Provider struct {
	ProviderID int    `json:"providerID"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type Ingredient struct {
	IngredientID int    `json:"ingredientID"`
	Name         string `json:"name"`
	ProviderID   int    `json:"providerID"`
}

type Dish struct {
	DishID      int    `json:"dishID"`
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Weight      int    `json:"weight"`
	CookingTime int    `json:"cookingTime"`
}

type Sale struct {
	DishID     int       `json:"dishID"`
	Percentage int       `json:"percentage"`
	BeginDate  time.Time `json:"beginDate"`
	EndDate    time.Time `json:"endDate"`
}

type Category struct {
	CategoryID  int    `json:"categoryID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
