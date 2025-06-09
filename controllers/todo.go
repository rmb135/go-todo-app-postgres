package controllers

import (
    "go-todo-app/config"
    "go-todo-app/models"
    "github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uint)
    var todos []models.Todo
    config.DB.Where("user_id = ?", userID).Find(&todos)
    return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uint)
    todo := new(models.Todo)
    if err := c.BodyParser(todo); err != nil {
        return err
    }
    todo.UserID = userID
    config.DB.Create(todo)
    return c.JSON(todo)
}
