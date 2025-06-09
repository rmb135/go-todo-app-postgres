package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret")

func JWTProtected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        tokenStr := c.Get("Authorization")
        if tokenStr == "" {
            return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
        }

        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
        }

        claims := token.Claims.(jwt.MapClaims)
        c.Locals("user_id", uint(claims["user_id"].(float64)))
        return c.Next()
    }
}
