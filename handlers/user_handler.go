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

func UpdateProfile(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    bio := c.FormValue("bio")
    profilePicture := c.FormValue("profile_picture")

    sql := `UPDATE users SET bio = $1, profile_picture = $2 WHERE user_id = $3`
    _, err := database.DB.Exec(context.Background(), sql, bio, profilePicture, userID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Profile updated successfully"})
}
