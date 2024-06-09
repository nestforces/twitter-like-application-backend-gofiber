package handlers

import (
    "context"
    "twitter-like-backend/database"
    "github.com/gofiber/fiber/v2"
)

func ChangeUsername(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    newUsername := c.FormValue("username")

    if newUsername == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "New username cannot be empty"})
    }

    sql := `UPDATE users SET username = $1 WHERE user_id = $2`
    _, err := database.DB.Exec(context.Background(), sql, newUsername, userID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update username"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Username updated successfully"})
}
