package routes

import (
    "github.com/gofiber/fiber/v2"
    "go-todo-app/controllers"
    "go-todo-app/middleware"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/register", controllers.Register)
    app.Post("/login", controllers.Login)

    api := app.Group("/api", middleware.JWTProtected())
    api.Get("/todos", controllers.GetTodos)
    api.Post("/todos", controllers.CreateTodo)
}
