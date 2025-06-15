package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang_crud/db"
	"golang_crud/db/models"

	"gorm.io/gorm" 
)

func main() {
	db.LoadEnv()
	db.ConnectDatabase()

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var users []models.User
		var products []models.Product
		var orders []models.Order

		// Insert 5 users
		for i := 1; i <= 5; i++ {
			user := models.User{
				Name:  fmt.Sprintf("User%d", i),
				Email: fmt.Sprintf("user%d@example.com", i),
			}
			users = append(users, user)
		}
		result := tx.Create(&users)
		if result.Error != nil {
			return result.Error
		}

		// Insert 20 products
		for i := 1; i <= 20; i++ {
			product := models.Product{
				Name:        fmt.Sprintf("Product%d", i),
				Description: fmt.Sprintf("Description for Product %d", i),
				Price:       float64(rand.Intn(5000) + 500),
				Quantity:    rand.Intn(50) + 10,
			}
			products = append(products, product)
		}
		result = tx.Create(&products)
		if result.Error != nil {
			return result.Error
		}

		// Insert 20 orders per user
		for _, user := range users {
			for j := 0; j < 20; j++ {
				randomProduct := products[rand.Intn(len(products))]
				order := models.Order{
					UserID:    user.ID,
					ProductID: randomProduct.ID,
					Qty:       rand.Intn(5) + 1,
					CreatedAt: time.Now(),
				}
				orders = append(orders, order)
			}
		}
		result = tx.Create(&orders)
		if result.Error != nil {
			return result.Error
		}

		// ✅ Return nil if all went well
		return nil
	})

	if err != nil {
		log.Fatal("Transaction failed:", err)
	} else {
		fmt.Println("✅ Transaction successful: 5 users, 20 products, and 100 orders created")
	}
}
