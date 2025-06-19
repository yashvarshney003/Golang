package db

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        AppConfig.Host, AppConfig.User, AppConfig.Password, AppConfig.Name, AppConfig.Port)

    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    DB = database
    log.Println("📦 Connected to Database!")
}
