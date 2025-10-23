package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var db *gorm.DB
var jwtKey = []byte("my_secret_key")

// функция генерации токена
func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	fmt.Println(token.SignedString(jwtKey))
	return token.SignedString(jwtKey)
}

// функция логина
func login(c *gin.Context) {
	var creds Credential
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	var flag bool
	flag = false
	for _, cred := range creditionals {
		if cred == creds {
			flag = true
			break
		}
	}
	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	token, err := generateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		fmt.Println(tokenString)
		tokenString = tokenString[7:]
		fmt.Println(tokenString)
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// все функции для всех путей получения данных
func allCRUDfuncs[T any](r *gin.Engine, path string, data *[]T, itemType reflect.Type) {

	r.GET(path, func(c *gin.Context) {
		db.Find(data)
		c.JSON(http.StatusOK, *data)
	})

	r.GET(path+"/:id", func(c *gin.Context) {
		id := c.Param("id")
		item := reflect.New(itemType).Interface()
		if err := db.First(&item, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		}
		c.JSON(http.StatusOK, &item)
	})

	r.DELETE(path+"/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&data, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
		return
	})

	r.POST(path, func(c *gin.Context) {
		newItem := reflect.New(itemType).Interface()
		if err := c.BindJSON(&newItem); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		db.Create(newItem)
		c.JSON(http.StatusCreated, newItem)
	})

	r.PUT(path+"/:id", func(c *gin.Context) {
		id := c.Param("id")
		updatedItem := reflect.New(itemType).Interface()
		if err := c.BindJSON(&updatedItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		if err := db.Model(reflect.New(itemType).Interface()).Where(fmt.Sprintf("%s_ID = ?", strings.ToLower(itemType.Name())), id).Updates(updatedItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		}
		c.JSON(http.StatusOK, updatedItem)
	})
}

// все функции для всех путей получения данных
func allPCRUDfuncs[T any](r *gin.Engine, path string, data *[]T, itemType reflect.Type) {
	protected := r.Group("/")
	protected.Use(authMiddleware())
	{
		protected.GET(path, func(c *gin.Context) {
			db.Find(data)
			c.JSON(http.StatusOK, *data)
		})

		protected.GET(path+"/:id", func(c *gin.Context) {
			id := c.Param("id")
			item := reflect.New(itemType).Elem()
			if err := db.First(&item, id).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
				return
			}
			c.JSON(http.StatusOK, &item)
		})

		protected.DELETE(path+"/:id", func(c *gin.Context) {
			id := c.Param("id")
			if err := db.Delete(&data, id).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
			return
		})

		protected.POST(path, func(c *gin.Context) {
			var item T
			if err := c.ShouldBindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data"})
				return
			}
			*data = append(*data, item)
			c.JSON(http.StatusCreated, gin.H{"message": "Item added"})
		})

		protected.PUT(path+"/:id", func(c *gin.Context) {
			id := c.Param("id")
			updatedItem := reflect.New(itemType).Interface()
			if err := c.BindJSON(&updatedItem); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
				return
			}
			if err := db.Model(reflect.New(itemType).Interface()).Where(fmt.Sprintf("%s_ID = ?", strings.ToLower(itemType.Name())), id).Updates(updatedItem).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
				return
			}
			c.JSON(http.StatusOK, updatedItem)
		})
	}
}

// функция получения ID структур
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

// функция запуска БД
func initDB() {
	dsn := "host=109.172.114.37 user=postgres password=123456 dbname=cafeAPI port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Market{})
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Provider{})
	db.AutoMigrate(&Ingredient{})
	db.AutoMigrate(&Dish{})
	db.AutoMigrate(&Sale{})
	db.AutoMigrate(&Category{})
}
