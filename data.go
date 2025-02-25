package main

import "time"

var customers = []Customer{
	{CustomerID: 1, FirstName: "Andrey", LastName: "Galanov", Email: "dawdwad@mail.ru", Phone: "89687889030", Bonuses: 99999999},
	{CustomerID: 2, FirstName: "Ilya", LastName: "Zmeev", Email: "qrmkqwelmr@mail.ru", Phone: "89128578934", Bonuses: 0},
	{CustomerID: 3, FirstName: "Egor", LastName: "Anikeev", Email: "ddwadwa@mail.ru", Phone: "87846789012", Bonuses: 100},
}

var markets = []Market{
	{MarketID: 1, Name: "Fresh Mart", Address: "123 Main St", StartHour: 8, EndHour: 22, Phone: "123-456-7890"},
	{MarketID: 2, Name: "Green Grocers", Address: "456 Oak Ave", StartHour: 7, EndHour: 21, Phone: "987-654-3210"},
	{MarketID: 3, Name: "Daily Foods", Address: "789 Pine Rd", StartHour: 9, EndHour: 20, Phone: "555-123-4567"},
}

var employees = []Employee{
	{EmployeeID: 1, FirstName: "John", LastName: "Doe", Passport: "A123456", Phone: "111-222-3333", PostID: 1},
	{EmployeeID: 2, FirstName: "Jane", LastName: "Smith", Passport: "B654321", Phone: "444-555-6666", PostID: 2},
	{EmployeeID: 3, FirstName: "Mike", LastName: "Johnson", Passport: "C789012", Phone: "777-888-9999", PostID: 3},
}

var posts = []Post{
	{PostID: 1, Name: "Cashier", Salary: 2500},
	{PostID: 2, Name: "Manager", Salary: 4000},
	{PostID: 3, Name: "Cleaner", Salary: 1800},
}

var orders = []Order{
	{OrderID: 1, DateTime: time.Now(), EmployeeID: 1, Closed: false, CustomerID: 1001},
	{OrderID: 2, DateTime: time.Now(), EmployeeID: 2, Closed: true, CustomerID: 1002},
	{OrderID: 3, DateTime: time.Now(), EmployeeID: 3, Closed: false, CustomerID: 1003},
}

var providers = []Provider{
	{ProviderID: 1, Name: "Food Supply Co.", Email: "contact@foodsupply.com", Phone: "101-202-3030"},
	{ProviderID: 2, Name: "Organic Goods", Email: "info@organicgoods.com", Phone: "404-505-6060"},
	{ProviderID: 3, Name: "Fresh Ingredients", Email: "support@freshing.com", Phone: "707-808-9090"},
}

var ingredients = []Ingredient{
	{IngredientID: 1, Name: "Tomato", ProviderID: 1},
	{IngredientID: 2, Name: "Cheese", ProviderID: 2},
	{IngredientID: 3, Name: "Lettuce", ProviderID: 3},
}

var dishes = []Dish{
	{DishID: 1, Name: "Pasta Carbonara", Photo: "pasta.jpg", Description: "Creamy Italian pasta", Cost: 15, Weight: 250, CookingTime: 20},
	{DishID: 2, Name: "Cheeseburger", Photo: "burger.jpg", Description: "Juicy beef burger with cheese", Cost: 10, Weight: 300, CookingTime: 15},
	{DishID: 3, Name: "Caesar Salad", Photo: "salad.jpg", Description: "Crisp romaine with dressing", Cost: 8, Weight: 200, CookingTime: 10},
}

var sales = []Sale{
	{SaleID: 1, Percentage: 10, BeginDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 7)},
	{SaleID: 2, Percentage: 15, BeginDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 10)},
	{SaleID: 3, Percentage: 6, BeginDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 5)},
}

var categories = []Category{
	{CategoryID: 1, Name: "Main Course", Description: "Hearty meals to fill you up"},
	{CategoryID: 2, Name: "Fast Food", Description: "Quick and tasty options"},
	{CategoryID: 3, Name: "Salads", Description: "Healthy and fresh dishes"},
}
