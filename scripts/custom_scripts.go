package main

import (
	"fmt"
	"log"
	"time"

	"golang_crud/db"
	"golang_crud/db/models"
)

func main() {
	db.LoadEnv()
	db.ConnectDatabase()

	var orders []models.Order

	// ✅ Join orders with users, filter where user name = 'User5', preload associations
	err := db.DB.
		Preload("User").
		Preload("Product").
		Joins("JOIN users ON users.id = orders.user_id").
		Where("users.name = ?", "User5").
		Find(&orders).Error

	if err != nil {
		log.Println("❌ Error fetching orders:", err)
		return
	}

	// ✅ Print order info
	for _, order := range orders {
		fmt.Println("------------")
		fmt.Printf("Order ID    : %d\n", order.ID)
		fmt.Printf("Qty         : %d\n", order.Qty)
		fmt.Printf("Created At  : %s\n", order.CreatedAt.Format(time.RFC822))
		fmt.Printf("User Name   : %s\n", order.User.Name)
		fmt.Printf("User Email  : %s\n", order.User.Email)
		fmt.Printf("Product Name: %s\n", order.Product.Name)
		fmt.Printf("Price       : %.2f\n", order.Product.Price)
	}
}
