package middlewares

import (
    "twitter-like-backend/utils"

    "github.com/gofiber/fiber/v2"
)

func AuthRequired() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if token == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
        }

        userID, err := utils.ParseJWT(token)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
        }

        c.Locals("user_id", userID)
        return c.Next()
    }
}
