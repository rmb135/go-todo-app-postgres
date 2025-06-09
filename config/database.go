package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "go-todo-app/models"
    "fmt"
    "os"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=todo_db port=5432 sslmode=disable")
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.User{}, &models.Todo{})
    DB = database
}
