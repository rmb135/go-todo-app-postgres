package main

import (
    "go-todo-app/config"
    "go-todo-app/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    config.ConnectDB()

    app := fiber.New()
    routes.SetupRoutes(app)

    app.Listen(":3000")
}
