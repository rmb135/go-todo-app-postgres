package controllers

import (
    "go-todo-app/config"
    "go-todo-app/models"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var jwtSecret = []byte("secret")

func Register(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return err
    }

    password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

    user := models.User{
        Username: data["username"],
        Password: string(password),
    }

    config.DB.Create(&user)
    return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return err
    }

    var user models.User
    config.DB.Where("username = ?", data["username"]).First(&user)

    if user.ID == 0 {
        return c.Status(404).SendString("User not found")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
        return c.Status(401).SendString("Incorrect password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    t, err := token.SignedString(jwtSecret)
    if err != nil {
        return err
    }

    return c.JSON(fiber.Map{"token": t})
}
