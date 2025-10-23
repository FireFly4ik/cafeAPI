package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Customer struct {
	CustomerID int    `gorm:"primaryKey" json:"customerID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Bonuses    int    `json:"bonuses"`
}

type Market struct {
	MarketID  int    `gorm:"primaryKey" json:"marketID"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	StartHour int    `json:"startHour"`
	EndHour   int    `json:"endHour"`
	Phone     string `json:"phone"`
}

type Employee struct {
	EmployeeID int    `gorm:"primaryKey" json:"employeeID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Passport   string `json:"passport"`
	Phone      string `json:"phone"`
	PostID     int    `json:"postID"`
}

type Post struct {
	PostID int    `gorm:"primaryKey" json:"postID"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type Order struct {
	OrderID    int       `gorm:"primaryKey" json:"orderID"`
	DateTime   time.Time `json:"dateTime"`
	EmployeeID int       `gorm:"primaryKey" json:"employeeID"`
	Closed     bool      `json:"closed"`
	CustomerID int       `gorm:"primaryKey" json:"customerID"`
}

type Provider struct {
	ProviderID int    `gorm:"primaryKey" json:"providerID"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type Ingredient struct {
	IngredientID int    `gorm:"primaryKey" json:"ingredientID"`
	Name         string `json:"name"`
	ProviderID   int    `json:"providerID"`
}

type Dish struct {
	DishID      int    `gorm:"primaryKey" json:"dishID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Cost        int    `json:"cost"`
	Weight      int    `json:"weight"`
	CookingTime int    `json:"cookingTime"`
}

type Sale struct {
	SaleID     int       `gorm:"primaryKey" json:"saleID"`
	Percentage int       `json:"percentage"`
	BeginDate  time.Time `json:"beginDate"`
	EndDate    time.Time `json:"endDate"`
}

type Category struct {
	CategoryID  int    `gorm:"primaryKey" json:"categoryID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
